package yandex_taxi_go

import (
	"context"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sinland/yandex-taxi-go/internal/models"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testClientID = "client-id"
	testAPIKey   = "api-key"
)

func TestClient_GetCarsList(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		testVehicle := models.Vehicle{
			Id:               "2111ade6gk054dfdb9iu8c8cc9460mks",
			Amenities:        []string{"wifi"},
			Brand:            "Mercedes-Benz",
			Callsign:         "123456789",
			Category:         []string{"econom"},
			Color:            "Черный",
			Model:            "E-klasse",
			Number:           "Т8654Т99",
			RegistrationCert: "123456789",
			Status:           "working",
			Vin:              "12345678909876543",
			Year:             2019,
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, testAPIKey, r.Header.Get(headerXAPIKey))
			require.Equal(t, testClientID, r.Header.Get(headerXCientID))

			w.WriteHeader(http.StatusOK)
			moqResult := models.CarsListResponse{
				Total:  1000,
				Offset: 0,
				Limit:  1000,
				Cars:   []models.Vehicle{testVehicle},
			}
			bytes, _ := json.Marshal(moqResult)
			_, err := w.Write(bytes)
			require.NoError(t, err)
		}))

		c := NewClient(ClientConfig{
			ClientID: testClientID,
			APIKey:   testAPIKey,
		}, WithAPIHost(server.URL))

		got, err := c.GetCarsList(ctx, GetCarsListArgs{
			Page:   0,
			Limit:  1000,
			ParkID: "park-id",
		})

		require.NoError(t, err)
		require.NotNil(t, got)

		require.Equal(t, 1000, got.Total)
		require.Equal(t, 1000, got.Limit)
		require.Equal(t, 0, got.Offset)
		require.Len(t, got.Cars, 1)
		require.Equal(t, got.Cars[0].Id, testVehicle.Id)
		require.Equal(t, got.Cars[0].Amenities, testVehicle.Amenities)
		require.Equal(t, got.Cars[0].Brand, testVehicle.Brand)
		require.Equal(t, got.Cars[0].Callsign, testVehicle.Callsign)
		require.Equal(t, got.Cars[0].Category, testVehicle.Category)
		require.Equal(t, got.Cars[0].Color, testVehicle.Color)
		require.Equal(t, got.Cars[0].Model, testVehicle.Model)
		require.Equal(t, got.Cars[0].Number, testVehicle.Number)
		require.Equal(t, got.Cars[0].RegistrationCert, testVehicle.RegistrationCert)
		require.Equal(t, got.Cars[0].Status, testVehicle.Status)
		require.Equal(t, got.Cars[0].Vin, testVehicle.Vin)
		require.Equal(t, got.Cars[0].Year, testVehicle.Year)
	})

	t.Run("failed request", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, testAPIKey, r.Header.Get(headerXAPIKey))
			require.Equal(t, testClientID, r.Header.Get(headerXCientID))
			//
			w.WriteHeader(http.StatusBadRequest)
			moqResult := models.ErrorResponse{
				Code:    "400",
				Message: "Bad request",
			}
			bytes, _ := json.Marshal(moqResult)
			_, err := w.Write(bytes)
			require.NoError(t, err)
		}))

		c := NewClient(ClientConfig{
			ClientID: testClientID,
			APIKey:   testAPIKey,
		}, WithAPIHost(server.URL))

		got, err := c.GetCarsList(ctx, GetCarsListArgs{
			Page:   0,
			Limit:  1000,
			ParkID: "park-id",
		})

		require.Error(t, err)
		require.Equal(t, "[400] Bad request (400)", err.Error())
		require.Nil(t, got)
	})
}

func TestClient_GetDriverProfiles(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := GetDriverProfilesArgs{
			Offset:    5,
			Limit:     1000,
			ParkId:    "park-id",
			QueryText: "some query text",
		}

		testPark := models.DriverProfilePark{
			Id:   gofakeit.UUID(),
			City: gofakeit.City(),
			Name: gofakeit.Company(),
		}

		testProfile := models.DriverProfile{
			Accounts: []models.DriverProfileAccount{{
				Id:           gofakeit.UUID(),
				Balance:      "1000.00",
				BalanceLimit: "50.00",
				Currency:     "RUB",
				Type:         "current",
			}},
			Car: &models.Vehicle{
				Id:               gofakeit.UUID(),
				Amenities:        []string{"conditioner"},
				Brand:            "Mercedes-Benz",
				Callsign:         "123456789",
				Category:         []string{"comfort_plus"},
				Color:            "Черный",
				Model:            "E-klasse",
				Number:           "Т8654Т99",
				RegistrationCert: "123456789",
				Status:           "working",
				Vin:              "12345678909876543",
				Year:             2019,
			},
			CurrentStatus: &models.DriverProfileCurrentStatus{
				Status:          "busy",
				StatusUpdatedAt: "2020-04-27T08:44:05.871+0000",
			},
			DriverProfile: &models.DriverProfileModel{
				Id:               gofakeit.UUID(),
				CheckMessage:     "great driver",
				Comment:          "great driver",
				CreatedDate:      "2020-04-23T13:08:05.552+0000",
				DriverLicense:    "AD12345",
				EmploymentType:   "selfemployed",
				FirstName:        "Ivan",
				HasContractIssue: true,
				LastName:         "Ivanov",
				MiddleName:       "Ivanovich",
				ParkId:           testPark.Id,
				Phones:           []string{"+79999999999"},
				WorkRuleId:       gofakeit.UUID(),
				WorkStatus:       "working",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, testAPIKey, r.Header.Get(headerXAPIKey))
			require.Equal(t, testClientID, r.Header.Get(headerXCientID))

			var req models.DriverProfilesRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			require.NoError(t, err)
			require.Equal(t, args.Limit, req.Limit)
			require.Equal(t, args.Offset, req.Offset)
			require.Equal(t, args.ParkId, req.Query.Park.Id)
			require.Equal(t, args.QueryText, req.Query.Text)

			w.WriteHeader(http.StatusOK)
			moqResult := models.DriverProfilesResponse{
				Total:          6543,
				Offset:         req.Offset,
				Limit:          req.Limit,
				DriverProfiles: []models.DriverProfile{testProfile},
				Parks:          []models.DriverProfilePark{testPark},
			}
			bytes, _ := json.Marshal(moqResult)
			_, err = w.Write(bytes)
			require.NoError(t, err)
		}))

		c := NewClient(ClientConfig{
			ClientID: testClientID,
			APIKey:   testAPIKey,
		}, WithAPIHost(server.URL))

		result, err := c.GetDriverProfiles(ctx, args)

		require.NoError(t, err)
		require.NotNil(t, result)

		require.Equal(t, 1, len(result.DriverProfiles))
		require.Equal(t, testProfile.DriverProfile.Id, result.DriverProfiles[0].Profile.Id)
		require.Equal(t, testProfile.DriverProfile.CheckMessage, result.DriverProfiles[0].Profile.CheckMessage)
		require.Equal(t, testProfile.DriverProfile.Comment, result.DriverProfiles[0].Profile.Comment)
		require.Equal(t, testProfile.DriverProfile.CreatedDate, result.DriverProfiles[0].Profile.CreatedDate)
		require.Equal(t, testProfile.DriverProfile.DriverLicense, result.DriverProfiles[0].Profile.DriverLicense)
		require.Equal(t, testProfile.DriverProfile.EmploymentType, result.DriverProfiles[0].Profile.EmploymentType)
		require.Equal(t, testProfile.DriverProfile.FirstName, result.DriverProfiles[0].Profile.FirstName)
		require.Equal(t, testProfile.DriverProfile.HasContractIssue, result.DriverProfiles[0].Profile.HasContractIssue)
		require.Equal(t, testProfile.DriverProfile.LastName, result.DriverProfiles[0].Profile.LastName)
		require.Equal(t, testProfile.DriverProfile.MiddleName, result.DriverProfiles[0].Profile.MiddleName)
		require.Equal(t, testProfile.DriverProfile.ParkId, result.DriverProfiles[0].Profile.ParkId)
		require.Equal(t, testProfile.DriverProfile.Phones, result.DriverProfiles[0].Profile.Phones)
		require.Equal(t, testProfile.DriverProfile.WorkRuleId, result.DriverProfiles[0].Profile.WorkRuleId)
		require.Equal(t, testProfile.DriverProfile.WorkStatus, result.DriverProfiles[0].Profile.WorkStatus)
		require.Equal(t, testProfile.CurrentStatus.Status, result.DriverProfiles[0].CurrentStatus.Status)
		require.Equal(t, testProfile.CurrentStatus.StatusUpdatedAt, result.DriverProfiles[0].CurrentStatus.StatusUpdatedAt)
		require.Equal(t, testProfile.Car.Id, result.DriverProfiles[0].Car.Id)
		require.Equal(t, testProfile.Car.Amenities, result.DriverProfiles[0].Car.Amenities)
		require.Equal(t, testProfile.Car.Brand, result.DriverProfiles[0].Car.Brand)
		require.Equal(t, testProfile.Car.Callsign, result.DriverProfiles[0].Car.Callsign)
		require.Equal(t, testProfile.Car.Category, result.DriverProfiles[0].Car.Category)
		require.Equal(t, testProfile.Car.Color, result.DriverProfiles[0].Car.Color)
		require.Equal(t, testProfile.Car.Model, result.DriverProfiles[0].Car.Model)
		require.Equal(t, testProfile.Car.Number, result.DriverProfiles[0].Car.Number)
		require.Equal(t, testProfile.Car.RegistrationCert, result.DriverProfiles[0].Car.RegistrationCert)
		require.Equal(t, testProfile.Car.Status, result.DriverProfiles[0].Car.Status)
		require.Equal(t, testProfile.Car.Vin, result.DriverProfiles[0].Car.Vin)
		require.Equal(t, testProfile.Car.Year, result.DriverProfiles[0].Car.Year)
		require.Equal(t, len(testProfile.Accounts), len(result.DriverProfiles[0].Accounts))
		require.Equal(t, testProfile.Accounts[0].Id, result.DriverProfiles[0].Accounts[0].Id)
		require.Equal(t, testProfile.Accounts[0].Balance, result.DriverProfiles[0].Accounts[0].Balance)
		require.Equal(t, testProfile.Accounts[0].BalanceLimit, result.DriverProfiles[0].Accounts[0].BalanceLimit)
		require.Equal(t, testProfile.Accounts[0].Currency, result.DriverProfiles[0].Accounts[0].Currency)
		require.Equal(t, testProfile.Accounts[0].Type, result.DriverProfiles[0].Accounts[0].Type)

		require.Equal(t, 1, len(result.Parks))
		require.Equal(t, testPark.Id, result.Parks[0].Id)
		require.Equal(t, testPark.City, result.Parks[0].City)
		require.Equal(t, testPark.Name, result.Parks[0].Name)
	})

	t.Run("failed request", func(t *testing.T) {
		t.Parallel()

		args := GetDriverProfilesArgs{
			Offset:    5,
			Limit:     1000,
			ParkId:    "park-id",
			QueryText: "some query text",
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, testAPIKey, r.Header.Get(headerXAPIKey))
			require.Equal(t, testClientID, r.Header.Get(headerXCientID))

			var req models.DriverProfilesRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			require.NoError(t, err)
			require.Equal(t, args.Limit, req.Limit)
			require.Equal(t, args.Offset, req.Offset)
			require.Equal(t, args.ParkId, req.Query.Park.Id)
			require.Equal(t, args.QueryText, req.Query.Text)

			w.WriteHeader(http.StatusBadRequest)
			moqResult := models.ErrorResponse{
				Code:    "400",
				Message: "Bad request",
			}
			bytes, _ := json.Marshal(moqResult)
			_, err = w.Write(bytes)
			require.NoError(t, err)
		}))

		c := NewClient(ClientConfig{
			ClientID: testClientID,
			APIKey:   testAPIKey,
		}, WithAPIHost(server.URL))

		result, err := c.GetDriverProfiles(ctx, args)

		require.Error(t, err)
		require.Equal(t, "[400] Bad request (400)", err.Error())
		require.Nil(t, result)
	})
}
