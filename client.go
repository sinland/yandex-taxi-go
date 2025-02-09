package yandex_taxi_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sinland/yandex-taxi-go/internal/models"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const (
	defaultApiHost   = "https://fleet-api.taxi.yandex.net"
	defaultPageLimit = 1000

	contentTypeJson = "application/json"

	headerContentType    = "Content-Type"
	headerAcceptLanguage = "Accept-Language"
	headerXAPIKey        = "X-API-Key"
	headerXCientID       = "X-Client-ID"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

// Client main api client
type Client struct {
	clientId   string
	apiKey     string
	apiHost    string
	httpClient httpClient
}

// NewClient constructor
func NewClient(cfg ClientConfig, opts ...ClientOption) *Client {
	c := &Client{
		apiKey:   cfg.APIKey,
		clientId: cfg.ClientID,
		apiHost:  defaultApiHost,
	}

	c.httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type ClientConfig struct {
	ClientID string
	APIKey   string
}

type ClientOption func(*Client)

func WithHttpClient(httpClient httpClient) func(client *Client) {
	return func(s *Client) {
		s.httpClient = httpClient
	}
}

func WithAPIHost(apiHost string) func(client *Client) {
	return func(s *Client) {
		s.apiHost = apiHost
	}
}

// GetCarsList Получение списка автомобилей
func (c *Client) GetCarsList(ctx context.Context, args GetCarsListArgs) (*GetCarsListResult, error) {
	reqUrl := fmt.Sprintf("%s/v1/parks/cars/list", c.apiHost)

	limit := args.Limit
	if limit == 0 {
		limit = defaultPageLimit
	}

	body, err := json.Marshal(models.CarsListRequest{
		Limit:  limit,
		Offset: args.Page * limit,
		Query: models.CarsListQuery{
			Park: models.CarsListQueryPark{
				Id: args.ParkID,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerContentType, contentTypeJson)
	req.Header.Set(headerXAPIKey, c.apiKey)
	req.Header.Set(headerXCientID, c.clientId)

	slog.DebugContext(ctx, "querying cars list", "url", reqUrl, "body", string(body))
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	slog.DebugContext(ctx, "query result", "status_code", res.StatusCode, "status", res.Status)
	if res.StatusCode != http.StatusOK {
		resData := models.ErrorResponse{}
		if err = json.NewDecoder(res.Body).Decode(&resData); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("[%d] %s (%s)", res.StatusCode, resData.Message, resData.Code)
	}

	var resData models.CarsListResponse
	if err = json.NewDecoder(res.Body).Decode(&resData); err != nil {
		return nil, err
	}

	result := &GetCarsListResult{
		Total:  resData.Total,
		Offset: resData.Offset,
		Limit:  resData.Limit,
		Cars:   make([]Vehicle, 0, len(resData.Cars)),
	}

	for i := range resData.Cars {
		result.Cars = append(result.Cars, Vehicle{
			Id:               resData.Cars[i].Id,
			Amenities:        resData.Cars[i].Amenities,
			Brand:            resData.Cars[i].Brand,
			Callsign:         resData.Cars[i].Callsign,
			Category:         resData.Cars[i].Category,
			Color:            resData.Cars[i].Color,
			Model:            resData.Cars[i].Model,
			Number:           resData.Cars[i].Number,
			RegistrationCert: resData.Cars[i].RegistrationCert,
			Status:           resData.Cars[i].Status,
			Vin:              resData.Cars[i].Vin,
			Year:             resData.Cars[i].Year,
		})
	}

	return result, nil
}

func (c *Client) GetDriverProfiles(ctx context.Context, args GetDriverProfilesArgs) (*GetDriverProfilesResult, error) {
	reqUrl := fmt.Sprintf("%s/v1/parks/driver-profiles/list", c.apiHost)

	limit := args.Limit
	if limit == 0 {
		limit = defaultPageLimit
	}

	body, err := json.Marshal(models.DriverProfilesRequest{
		Offset: args.Offset,
		Limit:  limit,
		Query: models.DriverProfilesListRequestQuery{
			Park: &models.DriverProfilesListRequestQueryPark{Id: args.ParkId},
			Text: args.QueryText,
		},
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerContentType, contentTypeJson)
	req.Header.Set(headerXAPIKey, c.apiKey)
	req.Header.Set(headerXCientID, c.clientId)
	req.Header.Set(headerAcceptLanguage, "ru")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resData := models.ErrorResponse{}
		if err = json.NewDecoder(res.Body).Decode(&resData); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("[%d] %s (%s)", res.StatusCode, resData.Message, resData.Code)
	}

	var resData models.DriverProfilesResponse
	if err = json.NewDecoder(res.Body).Decode(&resData); err != nil {
		return nil, err
	}

	result := &GetDriverProfilesResult{
		Total:          resData.Total,
		Offset:         resData.Offset,
		Limit:          resData.Limit,
		DriverProfiles: make([]DriverProfile, 0, len(resData.DriverProfiles)),
		Parks:          make([]DriverProfilePark, 0, len(resData.Parks)),
	}

	for i := range resData.Parks {
		result.Parks = append(result.Parks, DriverProfilePark{
			Id:   resData.Parks[i].Id,
			City: resData.Parks[i].City,
			Name: resData.Parks[i].Name,
		})
	}

	for i := range resData.DriverProfiles {
		profile := DriverProfile{
			Accounts: make([]DriverProfileAccount, 0, len(resData.DriverProfiles[i].Accounts)),
		}

		for j := range resData.DriverProfiles[i].Accounts {
			profile.Accounts = append(profile.Accounts, DriverProfileAccount{
				Id:           resData.DriverProfiles[i].Accounts[j].Id,
				Balance:      resData.DriverProfiles[i].Accounts[j].Balance,
				BalanceLimit: resData.DriverProfiles[i].Accounts[j].BalanceLimit,
				Currency:     resData.DriverProfiles[i].Accounts[j].Currency,
				Type:         resData.DriverProfiles[i].Accounts[j].Type,
			})
		}

		if resData.DriverProfiles[i].DriverProfile != nil {
			profile.Profile = &DriverProfileData{
				Id:           resData.DriverProfiles[i].DriverProfile.Id,
				CheckMessage: resData.DriverProfiles[i].DriverProfile.CheckMessage,
				Comment:      resData.DriverProfiles[i].DriverProfile.Comment,
				CreatedDate:  resData.DriverProfiles[i].DriverProfile.CreatedDate,
				DriverLicense: DriverLicense{
					IssueDate:        resData.DriverProfiles[i].DriverProfile.DriverLicense.IssueDate,
					ExpirationDate:   resData.DriverProfiles[i].DriverProfile.DriverLicense.ExpirationDate,
					Number:           resData.DriverProfiles[i].DriverProfile.DriverLicense.Number,
					NormalizedNumber: resData.DriverProfiles[i].DriverProfile.DriverLicense.NormalizedNumber,
					Country:          resData.DriverProfiles[i].DriverProfile.DriverLicense.Country,
					BirthDate:        resData.DriverProfiles[i].DriverProfile.DriverLicense.BirthDate,
				},
				EmploymentType:   resData.DriverProfiles[i].DriverProfile.EmploymentType,
				FirstName:        resData.DriverProfiles[i].DriverProfile.FirstName,
				HasContractIssue: resData.DriverProfiles[i].DriverProfile.HasContractIssue,
				LastName:         resData.DriverProfiles[i].DriverProfile.LastName,
				MiddleName:       resData.DriverProfiles[i].DriverProfile.MiddleName,
				ParkId:           resData.DriverProfiles[i].DriverProfile.ParkId,
				Phones:           resData.DriverProfiles[i].DriverProfile.Phones,
				WorkRuleId:       resData.DriverProfiles[i].DriverProfile.WorkRuleId,
				WorkStatus:       resData.DriverProfiles[i].DriverProfile.WorkStatus,
			}
		}

		if resData.DriverProfiles[i].Car != nil {
			profile.Car = &Vehicle{
				Id:               resData.DriverProfiles[i].Car.Id,
				Amenities:        resData.DriverProfiles[i].Car.Amenities,
				Brand:            resData.DriverProfiles[i].Car.Brand,
				Callsign:         resData.DriverProfiles[i].Car.Callsign,
				Category:         resData.DriverProfiles[i].Car.Category,
				Color:            resData.DriverProfiles[i].Car.Color,
				Model:            resData.DriverProfiles[i].Car.Model,
				Number:           resData.DriverProfiles[i].Car.Number,
				RegistrationCert: resData.DriverProfiles[i].Car.RegistrationCert,
				Status:           resData.DriverProfiles[i].Car.Status,
				Vin:              resData.DriverProfiles[i].Car.Vin,
				Year:             resData.DriverProfiles[i].Car.Year,
			}
		}

		if resData.DriverProfiles[i].CurrentStatus != nil {
			profile.CurrentStatus = &DriverProfileCurrentStatus{
				Status:          resData.DriverProfiles[i].CurrentStatus.Status,
				StatusUpdatedAt: resData.DriverProfiles[i].CurrentStatus.StatusUpdatedAt,
			}
		}

		result.DriverProfiles = append(result.DriverProfiles, profile)
	}

	return result, nil
}
