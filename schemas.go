package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. orderid is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. orderid is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "orderid": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "number"
                            },
                            "status": {
                    "description": "status",
                    "type": "string"
                },
                            "location": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
							"humidity": {
                                "description": "humidity",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"vibration": {
                                "description": "vibration",
                                "type": "number"
                            },
							"pressure": {
                                "description": "pressure",
                                "type": "number"
                            },
							"customername": {
                                "description": "customername",
                                "type": "string"
                            },
							"destination": {
                                "description": "destination",
                                "type": "string"
                            },
							"content": {
                                "description": "content",
                                "type": "string"
                            },
							"country": {
                                "description": "country",
                                "type": "string"
                            },
							"container": {
                                "description": "container",
                                "type": "string"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"health": {
                                "description": "health",
                                "type": "number"
                            }
							
							
                        },
                        "required": [
                            "orderid"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAsset function",
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset. Argument is a JSON encoded string containing only an orderid.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an orderid for use as an argument to read or delete.",
                        "properties": {
                            "orderid": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "number"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "nickname": {
                                "default": "SIMPLE",
                                "description": "The nickname of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "version"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. orderid is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an orderid for use as an argument to read or delete.",
                        "properties": {
                            "orderid": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "number"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                        "orderid": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "number"
                        },
                        "status": {
                    "description": "status",
                    "type": "string"
                },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        },
                        "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
							"humidity": {
                                "description": "humidity",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"vibration": {
                                "description": "vibration",
                                "type": "number"
                            },
							"pressure": {
                                "description": "pressure",
                                "type": "number"
                            },
							"customername": {
                                "description": "customername",
                                "type": "string"
                            },
							"destination": {
                                "description": "destination",
                                "type": "string"
                            },
							"content": {
                                "description": "content",
                                "type": "string"
                            },
							"country": {
                                "description": "country",
                                "type": "string"
                            },
							"container": {
                                "description": "container",
                                "type": "string"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"health": {
                                "description": "health",
                                "type": "number"
                            }
							
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetSamples": {
            "description": "Returns a string generated from the schema containing sample Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSamples function",
                    "enum": [
                        "readAssetSamples"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected sample data",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "readAssetSchemas": {
            "description": "Returns a string generated from the schema containing APIs and Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSchemas function",
                    "enum": [
                        "readAssetSchemas"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected schemas",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. orderid is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. orderid is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "orderid": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "number"
                            },
                            "status": {
                    "description": "status",
                    "type": "string"
                },
                            "location": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
							"humidity": {
                                "description": "humidity",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"vibration": {
                                "description": "vibration",
                                "type": "number"
                            },
							"pressure": {
                                "description": "pressure",
                                "type": "number"
                            },
							"customername": {
                                "description": "customername",
                                "type": "string"
                            },
							"destination": {
                                "description": "destination",
                                "type": "string"
                            },
							"content": {
                                "description": "content",
                                "type": "string"
                            },
							"country": {
                                "description": "country",
                                "type": "string"
                            },
							"container": {
                                "description": "container",
                                "type": "string"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"health": {
                                "description": "health",
                                "type": "number"
                            }
							
                        },
                        "required": [
                            "orderid"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "assetIDKey": {
            "description": "An object containing only an orderid for use as an argument to read or delete.",
            "properties": {
                "orderid": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "number"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an asset's state. orderid is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "orderid": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "number"
                },
                "status": {
                    "description": "status",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
							"humidity": {
                                "description": "humidity",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"vibration": {
                                "description": "vibration",
                                "type": "number"
                            },
							"pressure": {
                                "description": "pressure",
                                "type": "number"
                            },
							"customername": {
                                "description": "customername",
                                "type": "string"
                            },
							"destination": {
                                "description": "destination",
                                "type": "string"
                            },
							"content": {
                                "description": "content",
                                "type": "string"
                            },
							"country": {
                                "description": "country",
                                "type": "string"
                            },
							"container": {
                                "description": "container",
                                "type": "string"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"health": {
                                "description": "health",
                                "type": "number"
                            }
							
            },
            "required": [
                "orderid"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "SIMPLE",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of fields that constitute the complete asset state.",
            "properties": {
                "orderid": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "number"
                },
                "status": {
                    "description": "status",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
							"humidity": {
                                "description": "humidity",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"vibration": {
                                "description": "vibration",
                                "type": "number"
                            },
							"pressure": {
                                "description": "pressure",
                                "type": "number"
                            },
							"customername": {
                                "description": "customername",
                                "type": "string"
                            },
							"destination": {
                                "description": "destination",
                                "type": "string"
                            },
							"content": {
                                "description": "content",
                                "type": "string"
                            },
							"country": {
                                "description": "country",
                                "type": "string"
                            },
							"container": {
                                "description": "container",
                                "type": "string"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"health": {
                                "description": "health",
                                "type": "number"
                            }
							
            },
            "type": "object"
        }
    }
}`
