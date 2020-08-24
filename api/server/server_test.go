package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/YAITS/api/models"
	persistence "github.com/YAITS/api/persistence/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	server := getServer()
	startServer(server)

	t.Run("api", func(t *testing.T) {
		baseURL := fmt.Sprintf("http://%s/api", server.Addr)

		t.Run("root", func(t *testing.T) {
			url := fmt.Sprintf("%s/", baseURL)
			response, err := sendRequest(url, "GET", "")
			verifyResponse(t, response, err, http.StatusNotFound)
		})

		t.Run("HandleGETByID", func(t *testing.T) {
			url := fmt.Sprintf("%s/issue/1", baseURL)
			response, err := sendRequest(url, "GET", "")

			body, _ := ioutil.ReadAll(response.Body)
			var issueResponse models.IssueResponse
			_ = json.Unmarshal(body, &issueResponse)

			assert.Equal(t, persistence.MockIssueResponse, issueResponse, "issue response matches mock")
			verifyResponse(t, response, err, http.StatusOK)
		})

		t.Run("HandleGETByStatus", func(t *testing.T) {
			url := fmt.Sprintf("%s/issues/status?status=open", baseURL)
			response, err := sendRequest(url, "GET", "")

			body, _ := ioutil.ReadAll(response.Body)
			var issueResponse []models.IssueResponse
			_ = json.Unmarshal(body, &issueResponse)

			assert.Equal(t, []models.IssueResponse{persistence.MockIssueResponse}, issueResponse, "issue response matches mock")
			verifyResponse(t, response, err, http.StatusOK)
		})

		t.Run("HandleGETByPriority", func(t *testing.T) {
			url := fmt.Sprintf("%s/issues/priority?start=1", baseURL)
			response, err := sendRequest(url, "GET", "")

			body, _ := ioutil.ReadAll(response.Body)
			var issueResponse []models.IssueResponse
			_ = json.Unmarshal(body, &issueResponse)

			assert.Equal(t, []models.IssueResponse{persistence.MockIssueResponse}, issueResponse, "issue response matches mock")
			verifyResponse(t, response, err, http.StatusOK)
		})

		t.Run("HandlePOST", func(t *testing.T) {
			url := fmt.Sprintf("%s/issue", baseURL)
			requestBody := models.NewIssueRequest{
				Description: persistence.MockIssueResponse.Description,
				Priority:    persistence.MockIssueResponse.Priority,
				Summary:     persistence.MockIssueResponse.Summary,
			}

			requestBodyJSON, _ := json.Marshal(requestBody)

			response, err := sendRequest(url, "POST", string(requestBodyJSON))

			body, _ := ioutil.ReadAll(response.Body)
			var issueID int64
			_ = json.Unmarshal(body, &issueID)

			assert.Equal(t, issueID, int64(1), "a new issue id was returned")
			verifyResponse(t, response, err, http.StatusCreated)
		})

		t.Run("HandlePATCH", func(t *testing.T) {
			url := fmt.Sprintf("%s/issue/1", baseURL)
			requestBody := models.NewIssueRequest{
				Description: persistence.MockIssueResponse.Description,
				Priority:    persistence.MockIssueResponse.Priority,
				Summary:     persistence.MockIssueResponse.Summary,
			}

			requestBodyJSON, _ := json.Marshal(requestBody)

			response, err := sendRequest(url, "PATCH", string(requestBodyJSON))

			body, _ := ioutil.ReadAll(response.Body)
			var issueResponse models.IssueResponse
			_ = json.Unmarshal(body, &issueResponse)

			assert.Equal(t, persistence.MockIssueResponse, issueResponse, "a new issue id was returned")
			verifyResponse(t, response, err, http.StatusOK)
		})

		t.Run("HandleDELETE", func(t *testing.T) {
			url := fmt.Sprintf("%s/issue/1", baseURL)

			response, err := sendRequest(url, "DELETE", "")

			body, _ := ioutil.ReadAll(response.Body)
			var issueResponse models.IssueResponse
			_ = json.Unmarshal(body, &issueResponse)

			verifyResponse(t, response, err, http.StatusNoContent)
		})
	})
}

func startServer(s *http.Server) {
	go func() {
		_ = s.ListenAndServe()
	}()

	time.Sleep(100 * time.Millisecond)
}

func sendRequest(url string, method string, body string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func verifyResponse(t *testing.T, response *http.Response, err error, expectedStatus int) {
	if err != nil {
		t.Errorf("Error on response from server: %s", err)
	}

	if response == nil {
		t.Errorf("No response received from server")
		return
	}

	if response.StatusCode != expectedStatus {
		t.Errorf("Expected status %d, got %d", expectedStatus, response.StatusCode)
	}
}

func getServer() *http.Server {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	address := fmt.Sprintf("127.0.0.1:%d", port)

	logger := zap.NewNop().Sugar()
	storage := persistence.NewMockStorage()
	return NewServer(address, logger, storage)
}
