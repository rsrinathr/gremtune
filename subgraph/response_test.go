package subgraph

import (
	"log"
	"reflect"
	"testing"
)

/*
Dummy responses for mocking
*/

var dummySuccessfulResponse = []byte(`{    "requestId":"1d6d02bd-8e56-421d-9438-3bd6d0079ff1",    "status":{       "message":"",       "code":200,       "attributes":{ "@type":"g:Map", "@value":[     ]       }    },    "result":{       "data":{ "@type":"g:List", "@value":[    {       "@type":"tinker:graph",       "@value":{ "vertices":[    {  "@type":"g:Vertex",  "@value":{  "id":"a48f15ec-deae-55fb-b56b-77cb7dae9f3e",  "label":"Node::Product",  "properties":{     "name":[ {  "@type":"g:VertexProperty",  "@value":{     "id":{ "@type":"g:Int32", "@value":1254592157     },     "value":"AiringCanonicals",     "label":"name"  } }     ]    }  }    },    {  "@type":"g:Vertex",  "@value":{  "id":"c47eb7bb-d62c-5640-ab3b-1fb268a10fd9",  "label":"Node::Product",  "properties":{     "name":[ {  "@type":"g:VertexProperty",  "@value":{     "id":{ "@type":"g:Int32", "@value":-617607859     },     "value":"CAA",     "label":"name"  } }     ]  }  }    } ], "edges":[    {  "@type":"g:Edge",  "@value":{  "id":"04ba1ee1-a709-5660-c885-082602f0dae7",  "label":"describes",  "inVLabel":"Node::CatalogNode::Edit",  "outVLabel":"Node::TitleMetadata::Descriptor",  "inV":"11dada89-29b8-5792-bc55-49b86861c2cd",  "outV":"d6bbee47-73c1-54a4-bb83-1bb799b9b9ca",  "properties":{     "namespace":{ "@type":"g:Property", "@value":{  "key":"namespace",  "value":"54cc9992-cd93-5568-a65d-5e8c8f959edc" }     },     "bucketId":{ "@type":"g:Property", "@value":{  "key":"bucketId",  "value":{     "@type":"g:Int32",     "@value":2455  } }     },     "group":{ "@type":"g:Property", "@value":{  "key":"group",  "value":"f7ef2656-5577-3f2a-b921-4019eccf85a0" }     }  }  }    } ]      }    } ]       },       "meta":{ "@type":"g:Map", "@value":[     ]       }    } }`)

var dummyNeedAuthenticationResponse = []byte(`{"result":{},
 "requestId":"1d6d02bd-8e56-421d-9438-3bd6d0079ff1",
 "status":{"code":407,"attributes":{},"message":""}}`)

var dummyPartialResponse1 = []byte(`{    "requestId":"1d6d02bd-8e56-421d-9438-3bd6d0079ff1",    "status":{       "message":"",       "code":206,       "attributes":{ "@type":"g:Map", "@value":[     ]       }    },    "result":{       "data":{ "@type":"g:List", "@value":[    {       "@type":"tinker:graph",       "@value":{ "vertices":[    {  "@type":"g:Vertex",  "@value":{  "id":"a48f15ec-deae-55fb-b56b-77cb7dae9f3e",  "label":"Node::Product",  "properties":{     "name":[ {  "@type":"g:VertexProperty",  "@value":{     "id":{ "@type":"g:Int32", "@value":1254592157     },     "value":"AiringCanonicals",     "label":"name"  } }     ]    }  }    },    {  "@type":"g:Vertex",  "@value":{  "id":"c47eb7bb-d62c-5640-ab3b-1fb268a10fd9",  "label":"Node::Product",  "properties":{     "name":[ {  "@type":"g:VertexProperty",  "@value":{     "id":{ "@type":"g:Int32", "@value":-617607859     },     "value":"CAA",     "label":"name"  } }     ]  }  }    } ], "edges":[    {  "@type":"g:Edge",  "@value":{  "id":"04ba1ee1-a709-5660-c885-082602f0dae7",  "label":"describes",  "inVLabel":"Node::CatalogNode::Edit",  "outVLabel":"Node::TitleMetadata::Descriptor",  "inV":"11dada89-29b8-5792-bc55-49b86861c2cd",  "outV":"d6bbee47-73c1-54a4-bb83-1bb799b9b9ca",  "properties":{     "namespace":{ "@type":"g:Property", "@value":{  "key":"namespace",  "value":"54cc9992-cd93-5568-a65d-5e8c8f959edc" }     },     "bucketId":{ "@type":"g:Property", "@value":{  "key":"bucketId",  "value":{     "@type":"g:Int32",     "@value":2455  } }     },     "group":{ "@type":"g:Property", "@value":{  "key":"group",  "value":"f7ef2656-5577-3f2a-b921-4019eccf85a0" }     }  }  }    } ]      }    } ]       },       "meta":{ "@type":"g:Map", "@value":[     ]       }    } }`)

var dummyPartialResponse2 = []byte(`{    "requestId":"1d6d02bd-8e56-421d-9438-3bd6d0079ff1",    "status":{       "message":"",       "code":200,       "attributes":{ "@type":"g:Map", "@value":[     ]       }    },    "result":{       "data":{ "@type":"g:List", "@value":[    {       "@type":"tinker:graph",       "@value":{ "vertices":[    {  "@type":"g:Vertex",  "@value":{  "id":"a48f15ec-deae-55fb-b56b-77cb7dae9f3e",  "label":"Node::Product",  "properties":{     "name":[ {  "@type":"g:VertexProperty",  "@value":{     "id":{ "@type":"g:Int32", "@value":1254592157     },     "value":"AiringCanonicals",     "label":"name"  } }     ]    }  }    },    {  "@type":"g:Vertex",  "@value":{  "id":"c47eb7bb-d62c-5640-ab3b-1fb268a10fd9",  "label":"Node::Product",  "properties":{     "name":[ {  "@type":"g:VertexProperty",  "@value":{     "id":{ "@type":"g:Int32", "@value":-617607859     },     "value":"CAA",     "label":"name"  } }     ]  }  }    } ], "edges":[    {  "@type":"g:Edge",  "@value":{  "id":"04ba1ee1-a709-5660-c885-082602f0dae7",  "label":"describes",  "inVLabel":"Node::CatalogNode::Edit",  "outVLabel":"Node::TitleMetadata::Descriptor",  "inV":"11dada89-29b8-5792-bc55-49b86861c2cd",  "outV":"d6bbee47-73c1-54a4-bb83-1bb799b9b9ca",  "properties":{     "namespace":{ "@type":"g:Property", "@value":{  "key":"namespace",  "value":"54cc9992-cd93-5568-a65d-5e8c8f959edc" }     },     "bucketId":{ "@type":"g:Property", "@value":{  "key":"bucketId",  "value":{     "@type":"g:Int32",     "@value":2455  } }     },     "group":{ "@type":"g:Property", "@value":{  "key":"group",  "value":"f7ef2656-5577-3f2a-b921-4019eccf85a0" }     }  }  }    } ]      }    } ]       },       "meta":{ "@type":"g:Map", "@value":[     ]       }    } }`)

var dummySuccessfulResponseMarshalled = Response{
	RequestID: "1d6d02bd-8e56-421d-9438-3bd6d0079ff1",
	Status:    Status{Code: 200},
	Result:    Result{Data: NeptuneSubGraph{}},
}

var dummyNeedAuthenticationResponseMarshalled = Response{
	RequestID: "1d6d02bd-8e56-421d-9438-3bd6d0079ff1",
	Status:    Status{Code: 407},
	Result:    Result{Data: NeptuneSubGraph{}},
}

var dummyPartialResponse1Marshalled = Response{
	RequestID: "1d6d02bd-8e56-421d-9438-3bd6d0079ff1",
	Status:    Status{Code: 206}, // Code 206 indicates that the response is not the terminating response in a sequence of responses
	Result:    Result{Data: NeptuneSubGraph{}},
}

var dummyPartialResponse2Marshalled = Response{
	RequestID: "1d6d02bd-8e56-421d-9438-3bd6d0079ff1",
	Status:    Status{Code: 200},
	Result:    Result{Data: NeptuneSubGraph{}},
}

// TestResponseHandling tests the overall response handling mechanism of gremtune
func TestResponseHandling(t *testing.T) {
	c := newClient()

	c.handleResponse(dummySuccessfulResponse)

	var expected []*Response
	expected = append(expected, &dummySuccessfulResponseMarshalled)

	r, _ := c.retrieveResponse(dummySuccessfulResponseMarshalled.RequestID)
	if reflect.TypeOf(expected).String() != reflect.TypeOf(r).String() {
		t.Error("Expected data type does not match actual.")
	}
}

func TestResponseAuthHandling(t *testing.T) {
	c := newClient()
	ws := new(Ws)
	ws.auth = &auth{username: "test", password: "test"}
	c.conn = ws
	c.handleResponse(dummyNeedAuthenticationResponse)

	req, err := prepareAuthRequest(dummyNeedAuthenticationResponseMarshalled.RequestID, "test", "test")
	if err != nil {
		return
	}

	sampleAuthRequest, err := packageRequest(req)
	if err != nil {
		log.Println(err)
		return
	}

	c.dispatchRequest(sampleAuthRequest)
	authRequest := <-c.requests //Simulate that client send auth challenge to server
	if !reflect.DeepEqual(authRequest, sampleAuthRequest) {
		t.Error("Expected data type does not match actual.")
	}

	c.handleResponse(dummySuccessfulResponse) //If authentication is successful the server returns the origin petition

	var expectedSuccessful []*Response
	expectedSuccessful = append(expectedSuccessful, &dummySuccessfulResponseMarshalled)

	r, _ := c.retrieveResponse(dummySuccessfulResponseMarshalled.RequestID)
	if reflect.TypeOf(expectedSuccessful).String() != reflect.TypeOf(r).String() {
		t.Error("Expected data type does not match actual.")
	}
}

// TestResponseSortingSingleResponse tests the ability for sortResponse to save a response received from Gremlin Server
func TestResponseSortingSingleResponse(t *testing.T) {

	c := newClient()

	c.saveResponse(&dummySuccessfulResponseMarshalled, nil)

	var expected []*Response
	expected = append(expected, &dummySuccessfulResponseMarshalled)

	result, _ := c.results.Load(dummySuccessfulResponseMarshalled.RequestID)
	if reflect.DeepEqual(result.([]*Response), expected) != true {
		t.Fail()
	}
}

// TestResponseSortingMultipleResponse tests the ability for the sortResponse function to categorize and group responses that are sent in a stream
func TestResponseSortingMultipleResponse(t *testing.T) {

	c := newClient()

	c.saveResponse(&dummyPartialResponse1Marshalled, nil)
	c.saveResponse(&dummyPartialResponse2Marshalled, nil)

	var expected []*Response
	expected = append(expected, &dummyPartialResponse1Marshalled)
	expected = append(expected, &dummyPartialResponse2Marshalled)

	results, _ := c.results.Load(dummyPartialResponse1Marshalled.RequestID)
	if reflect.DeepEqual(results.([]*Response), expected) != true {
		t.Fail()
	}
}

// TestResponseRetrieval tests the ability for a requester to retrieve the response for a specified requestId generated when sending the request
func TestResponseRetrieval(t *testing.T) {
	c := newClient()

	c.saveResponse(&dummyPartialResponse1Marshalled, nil)
	c.saveResponse(&dummyPartialResponse2Marshalled, nil)

	resp, _ := c.retrieveResponse(dummyPartialResponse1Marshalled.RequestID)

	var expected []*Response
	expected = append(expected, &dummyPartialResponse1Marshalled)
	expected = append(expected, &dummyPartialResponse2Marshalled)

	if reflect.DeepEqual(resp, expected) != true {
		t.Fail()
	}
}

// TestResponseDeletion tests the ability for a requester to clean up after retrieving a response after delivery to a client
func TestResponseDeletion(t *testing.T) {
	c := newClient()

	c.saveResponse(&dummyPartialResponse1Marshalled, nil)
	c.saveResponse(&dummyPartialResponse2Marshalled, nil)

	c.deleteResponse(dummyPartialResponse1Marshalled.RequestID)

	if _, ok := c.results.Load(dummyPartialResponse1Marshalled.RequestID); ok {
		t.Fail()
	}
}

var codes = []struct {
	code int
}{
	{200},
	{204},
	{206},
	{401},
	{407},
	{498},
	{499},
	{500},
	{597},
	{598},
	{599},
	{3434}, // Testing unknown error code
}

// Tests detection of errors and if an error is generated for a specific error code
func TestResponseErrorDetection(t *testing.T) {
	for _, co := range codes {
		dummyResponse := Response{
			RequestID: "",
			Status:    Status{Code: co.code},
			Result:    Result{},
		}
		err := dummyResponse.detectError()
		switch {
		case co.code == 200:
			if err != nil {
				t.Log("Successful response returned error.")
			}
		case co.code == 204:
			if err != nil {
				t.Log("Successful response returned error.")
			}
		case co.code == 206:
			if err != nil {
				t.Log("Successful response returned error.")
			}
		default:
			if err == nil {
				t.Log("Unsuccessful response did not return error.")
			}
		}
	}
}
