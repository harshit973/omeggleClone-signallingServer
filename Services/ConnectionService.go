package Services

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"net/http"
	"omeggleClone-signallingServer/DTO"
	"omeggleClone-signallingServer/Exceptions"
	"omeggleClone-signallingServer/Repository"
)

func RequestOfferService(r *http.Request) (*string, *Exceptions.ApplicationException) {
	var payload DTO.RequestPayload
	err := payload.BuildPayloadFromRequest(r.Body)
	if err != nil {
		return nil, Exceptions.NewApplicationException(400, "Invalid payload", nil)
	}
	strangerConnectionID := payload.ConnectionID

	var connectionConfig DTO.ConnectionConfig
	connectionConfig.NewConnectionConfig(r.Context())
	client := connectionConfig.GetApiGatewayClient()
	if strangerConnectionID == nil {
		connectionID, err := Repository.FindARandomConnectionExcept(connectionConfig.ConnectionId)
		if err != nil {
			return nil, err
		}
		strangerConnectionID = connectionID
	}

	responseMessage := fmt.Sprintf(`{"message":"%v","strangerID":"%v"}`, connectionConfig.ConnectionId, payload.Message)
	_, err = client.PostToConnection(r.Context(), &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionConfig.ConnectionId),
		Data:         []byte(responseMessage),
	})
	if err != nil {
		return nil, Exceptions.NewApplicationException(400, "Unable to post to connection", &err)
	}
	return strangerConnectionID, nil
}

func CreateConnectionService(r *http.Request) *Exceptions.ApplicationException {
	currentConnectionID := r.Context().Value("connectionId")
	if currentConnectionID == nil {
		return Exceptions.NewApplicationException(404, "Invalid connectionID", nil)
	}
	err := Repository.CreateConnection(currentConnectionID.(string))
	if err != nil {
		return err
	}
	return nil
}

func DeactivateConnectionService(r *http.Request) *Exceptions.ApplicationException {
	currentConnectionID := r.Context().Value("connectionId")
	if currentConnectionID == nil {
		return Exceptions.NewApplicationException(404, "Invalid connectionID", nil)
	}
	_ = Repository.DeactivateConnection(currentConnectionID.(string))
	return nil
}
