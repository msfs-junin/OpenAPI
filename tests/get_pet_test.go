package integration

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
)

func TestCallGetPet(t *testing.T) {

	Convey("Given a the system is up and running", t, func() {
		mockHost := os.Getenv("MOCK_HOST")
		if mockHost == "" {
			mockHost = "localhost:1080"
		}

		Convey("When the user wants to get info about a pet", func() {
			req, err := http.NewRequest(http.MethodGet, "http://"+mockHost+"/pet/1", nil)
			So(err, ShouldEqual, nil)

			req.Header.Add("api_key", "xyz")

			client := &http.Client{}
			resp, err := client.Do(req)
			So(err, ShouldEqual, nil)

			defer func() {
				_ = resp.Body.Close()
			}()

			Convey("He gets the info about the pet", func() {
				body, err := ioutil.ReadAll(resp.Body)
				So(err, ShouldEqual, nil)

				expectedBodyJSON := `{
					"id": 1,
					"name": "Tom",
					"photoUrls": [
						"https://via.placeholder.com/150"
					]
				}`

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				require.JSONEq(t, expectedBodyJSON, string(body))
			})

		})

	})

}
