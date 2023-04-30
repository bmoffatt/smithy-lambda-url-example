// Code generated by smithy-go-codegen DO NOT EDIT.


package model

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"encoding/json"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	smithy "github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithyio "github.com/aws/smithy-go/io"
	smithytime "github.com/aws/smithy-go/time"
)

type awsRestjson1_deserializeOpWave struct {
}

func (*awsRestjson1_deserializeOpWave) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpWave) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil { return out, metadata, err }
	
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}
	
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorWave(response, &metadata)
	}
	output := &WaveOutput{}
	out.Result = output
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(response.Body, ringBuffer)
	
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}
	
	err = awsRestjson1_deserializeOpDocumentWaveOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}
	
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorWave(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())
	
	errorCode := "UnknownError"
	errorMessage := errorCode
	
	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 { errorCode = restjson.SanitizeErrorCode(headerCode) }
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 { errorCode = restjson.SanitizeErrorCode(jsonCode) }
	if len(message) != 0 { errorMessage = message }
	
	switch {
		default:
			genericError := &smithy.GenericAPIError{
				Code: errorCode,
				Message: errorMessage,
			}
			return genericError
		
	}
}

func awsRestjson1_deserializeOpDocumentWaveOutput(v **WaveOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *WaveOutput
	if *v == nil {
		sv = &WaveOutput{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "Text":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Text = ptr.String(jtv)
				}
			
			case "Time":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected Timestamp to be of type string, got %T instead", value)
					}
					t, err := smithytime.ParseDateTime(jtv)
					if err != nil { return err }
					sv.Time = ptr.Time(t)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}
