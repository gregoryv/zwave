package zwave

import "fmt"

// Headers
const (
	SOF byte = 0x01 // Start Of Frame
	ACK byte = 0x06 // Message Ack
	NAK byte = 0x15 // Message NAK
	CAN byte = 0x18 // Cancel - Resend request
)

type Named struct {
	Type   string
	Values map[byte]string
}

// lookup tries to find a named constant
func (n *Named) Get(key byte) string {
	name, found := n.Values[key]
	if !found {
		name = fmt.Sprintf("Unknown %s: %v", n.Type, key)
	}
	return name
}

var NamedHeader = &Named{
	"Header",
	map[byte]string{
		SOF: "SOF",
		ACK: "ACK",
		NAK: "NAK",
		CAN: "CAN",
	},
}

// Message types
const (
	REQUEST  byte = 0x00
	RESPONSE byte = 0x01
)

var NamedMessageType = &Named{
	"Message Type",
	map[byte]string{
		REQUEST:  "Request",
		RESPONSE: "Response",
	},
}

// Defined functions, taken from github code
const (
	FUNCTION_NONE                              = 0x0
	FUNCTION_DISCOVERY_NODES                   = 0x2
	FUNCTION_SERIAL_API_APPL_NODE_INFORMATION  = 0x3
	FUNCTION_APPLICATION_COMMAND_HANDLER       = 0x4
	FUNCTION_GET_CONTROLLER_CAPABILITIES       = 0x5
	FUNCTION_SERIAL_API_SET_TIMEOUTS           = 0x6
	FUNCTION_SERIAL_GET_CAPABILITIES           = 0x7
	FUNCTION_SERIAL_API_SOFT_RESET             = 0x8
	FUNCTION_SET_RF_RECEIVE_MODE               = 0x10
	FUNCTION_SET_SLEEP_MODE                    = 0x11
	FUNCTION_SEND_NODE_INFORMATION             = 0x12
	FUNCTION_SEND_DATA                         = 0x13
	FUNCTION_SEND_DATA_MULTI                   = 0x14
	FUNCTION_GET_VERSION                       = 0x15
	FUNCTION_SEND_DATA_ABORT                   = 0x16
	FUNCTION_RF_POWER_LEVEL_SET                = 0x17
	FUNCTION_SEND_DATA_META                    = 0x18
	FUNCTION_MEMORY_GET_ID                     = 0x20
	FUNCTION_MEMORY_GET_BYTE                   = 0x21
	FUNCTION_MEMORY_PUT_BYTE                   = 0x22
	FUNCTION_MEMORY_GET_BUFFER                 = 0x23
	FUNCTION_MEMORY_PUT_BUFFER                 = 0x24
	FUNCTION_READ_MEMORY                       = 0x25
	FUNCTION_CLOCK_SET                         = 0x30
	FUNCTION_CLOCK_GET                         = 0x31
	FUNCTION_CLOCK_COMPARE                     = 0x32
	FUNCTION_RTC_TIMER_CREATE                  = 0x33
	FUNCTION_RTC_TIMER_READ                    = 0x34
	FUNCTION_RTC_TIMER_DELETE                  = 0x35
	FUNCTION_RTC_TIMER_CALL                    = 0x36
	FUNCTION_GET_NODE_PROTOCOL_INFO            = 0x41
	FUNCTION_SET_DEFAULT                       = 0x42
	FUNCTION_REPLICATION_COMMAND_COMPLETE      = 0x44
	FUNCTION_REPLICATION_SEND_DATA             = 0x45
	FUNCTION_ASSIGN_RETURN_ROUTE               = 0x46
	FUNCTION_DELETE_RETURN_ROUTE               = 0x47
	FUNCTION_REQUEST_NODE_NEIGHBOR_UPDATE      = 0x48
	FUNCTION_APPLICATION_UPDATE                = 0x49
	FUNCTION_ADD_NODE_TO_NETWORK               = 0x4a
	FUNCTION_REMOVE_NODE_FROM_NETWORK          = 0x4b
	FUNCTION_CREATE_NEW_PRIMARY                = 0x4c
	FUNCTION_CONTROLLER_CHANGE                 = 0x4d
	FUNCTION_SET_LEARN_MODE                    = 0x50
	FUNCTION_ASSIGN_SUC_RETURN_ROUTE           = 0x51
	FUNCTION_ENABLE_SUC                        = 0x52
	FUNCTION_REQUEST_NETWORK_UPDATE            = 0x53
	FUNCTION_SET_SUC_NODE_ID                   = 0x54
	FUNCTION_DELETE_SUC_RETURN_ROUTE           = 0x55
	FUNCTION_GET_SUC_NODE_ID                   = 0x56
	FUNCTION_SEND_SUC_ID                       = 0x57
	FUNCTION_REDISCOVERY_NEEDED                = 0x59
	FUNCTION_REQUEST_NODE_INFO                 = 0x60
	FUNCTION_REMOVE_FAILED_NODE_ID             = 0x61
	FUNCTION_IS_FAILED_NODE                    = 0x62
	FUNCTION_REPLACE_FAILED_NODE               = 0x63
	FUNCTION_TIMER_START                       = 0x70
	FUNCTION_TIMER_RESTART                     = 0x71
	FUNCTION_TIMER_CANCEL                      = 0x72
	FUNCTION_TIMER_CALL                        = 0x73
	FUNCTION_GET_ROUTING_TABLE_LINE            = 0x80
	FUNCTION_GET_TX_COUNTER                    = 0x81
	FUNCTION_RESET_TX_COUNTER                  = 0x82
	FUNCTION_STORE_NODE_INFO                   = 0x83
	FUNCTION_STORE_HOME_ID                     = 0x84
	FUNCTION_LOCK_ROUTE_RESPONSE               = 0x90
	FUNCTION_SEND_DATA_ROUTE_DEMO              = 0x91
	FUNCTION_SERIAL_API_TEST                   = 0x95
	FUNCTION_SERIAL_API_SLAVE_NODE_INFO        = 0xa0
	FUNCTION_APPLICATION_SLAVE_COMMAND_HANDLER = 0xa1
	FUNCTION_SEND_SLAVE_NODE_INFO              = 0xa2
	FUNCTION_SEND_SLAVE_DATA                   = 0xa3
	FUNCTION_SET_SLAVE_LEARN_MODE              = 0xa4
	FUNCTION_GET_VIRTUAL_NODES                 = 0xa5
	FUNCTION_IS_VIRTUAL_NODE                   = 0xa6
	FUNCTION_SET_PROMISCUOUS_MODE              = 0xd0
)

var NamedFunction = &Named{
	"Function",
	map[byte]string{
		FUNCTION_NONE:                              "None",
		FUNCTION_DISCOVERY_NODES:                   "Discovery Nodes",
		FUNCTION_SERIAL_API_APPL_NODE_INFORMATION:  "Serial API Application Node Information",
		FUNCTION_APPLICATION_COMMAND_HANDLER:       "Application Command Handler",
		FUNCTION_GET_CONTROLLER_CAPABILITIES:       "Get Controller Capabilities",
		FUNCTION_SERIAL_API_SET_TIMEOUTS:           "Serial API Set Timeouts",
		FUNCTION_SERIAL_GET_CAPABILITIES:           "Serial Get Capabilities",
		FUNCTION_SERIAL_API_SOFT_RESET:             "Serial API Soft Reset",
		FUNCTION_SET_RF_RECEIVE_MODE:               "Set RF Receive Mode",
		FUNCTION_SET_SLEEP_MODE:                    "Set Sleep Mode",
		FUNCTION_SEND_NODE_INFORMATION:             "Send Node Information",
		FUNCTION_SEND_DATA:                         "Send Data",
		FUNCTION_SEND_DATA_MULTI:                   "Send Data Multi",
		FUNCTION_GET_VERSION:                       "Get Version",
		FUNCTION_SEND_DATA_ABORT:                   "Send Data Abort",
		FUNCTION_RF_POWER_LEVEL_SET:                "RF Power Level Set",
		FUNCTION_SEND_DATA_META:                    "Send Data Meta",
		FUNCTION_MEMORY_GET_ID:                     "Memory Get Id",
		FUNCTION_MEMORY_GET_BYTE:                   "Memory Get Byte",
		FUNCTION_MEMORY_PUT_BYTE:                   "Memory Put Byte",
		FUNCTION_MEMORY_GET_BUFFER:                 "Memory Get Buffer",
		FUNCTION_MEMORY_PUT_BUFFER:                 "Memory Put Buffer",
		FUNCTION_READ_MEMORY:                       "Read Memory",
		FUNCTION_CLOCK_SET:                         "Clock Set",
		FUNCTION_CLOCK_GET:                         "Clock Get",
		FUNCTION_CLOCK_COMPARE:                     "Clock Compare",
		FUNCTION_RTC_TIMER_CREATE:                  "RTC Timer Create",
		FUNCTION_RTC_TIMER_READ:                    "RTC Timer Read",
		FUNCTION_RTC_TIMER_DELETE:                  "RTC Timer Delete",
		FUNCTION_RTC_TIMER_CALL:                    "RTC Timer Call",
		FUNCTION_GET_NODE_PROTOCOL_INFO:            "Get Node Protocol Info",
		FUNCTION_SET_DEFAULT:                       "Set Default",
		FUNCTION_REPLICATION_COMMAND_COMPLETE:      "Replication Command Complete",
		FUNCTION_REPLICATION_SEND_DATA:             "Replication Send Data",
		FUNCTION_ASSIGN_RETURN_ROUTE:               "Assign Return Route",
		FUNCTION_DELETE_RETURN_ROUTE:               "Delete Return Route",
		FUNCTION_REQUEST_NODE_NEIGHBOR_UPDATE:      "Request Node Neighbor Update",
		FUNCTION_APPLICATION_UPDATE:                "Application Update",
		FUNCTION_ADD_NODE_TO_NETWORK:               "Add Node To Network",
		FUNCTION_REMOVE_NODE_FROM_NETWORK:          "Remove Node From Network",
		FUNCTION_CREATE_NEW_PRIMARY:                "Create New Primary",
		FUNCTION_CONTROLLER_CHANGE:                 "Controller Change",
		FUNCTION_SET_LEARN_MODE:                    "Set Learn Mode",
		FUNCTION_ASSIGN_SUC_RETURN_ROUTE:           "Assign Static Update Controller Return Route",
		FUNCTION_ENABLE_SUC:                        "Enable Static Update Controller",
		FUNCTION_REQUEST_NETWORK_UPDATE:            "Request Network Update",
		FUNCTION_SET_SUC_NODE_ID:                   "Set Static Update Controller Node Id",
		FUNCTION_DELETE_SUC_RETURN_ROUTE:           "Delete Static Update Controller Return Route",
		FUNCTION_GET_SUC_NODE_ID:                   "Get Static Update Controller Node Id",
		FUNCTION_SEND_SUC_ID:                       "Send Static Update Controller Id",
		FUNCTION_REDISCOVERY_NEEDED:                "Rediscovery Needed",
		FUNCTION_REQUEST_NODE_INFO:                 "Request Node Info",
		FUNCTION_REMOVE_FAILED_NODE_ID:             "Remove Failed Node Id",
		FUNCTION_IS_FAILED_NODE:                    "Is Failed Node",
		FUNCTION_REPLACE_FAILED_NODE:               "Replace Failed Node",
		FUNCTION_TIMER_START:                       "Timer Start",
		FUNCTION_TIMER_RESTART:                     "Timer Restart",
		FUNCTION_TIMER_CANCEL:                      "Timer Cancel",
		FUNCTION_TIMER_CALL:                        "Timer Call",
		FUNCTION_GET_ROUTING_TABLE_LINE:            "Get Routing Table Line",
		FUNCTION_GET_TX_COUNTER:                    "Get TX Counter",
		FUNCTION_RESET_TX_COUNTER:                  "Reset TX Counter",
		FUNCTION_STORE_NODE_INFO:                   "Store Node Info",
		FUNCTION_STORE_HOME_ID:                     "Store Home Id",
		FUNCTION_LOCK_ROUTE_RESPONSE:               "Lock Route Response",
		FUNCTION_SEND_DATA_ROUTE_DEMO:              "Send Data Route Demo",
		FUNCTION_SERIAL_API_TEST:                   "Serial API Test",
		FUNCTION_SERIAL_API_SLAVE_NODE_INFO:        "Serial API Slave Node Info",
		FUNCTION_APPLICATION_SLAVE_COMMAND_HANDLER: "Application Slave Command Handler",
		FUNCTION_SEND_SLAVE_NODE_INFO:              "Send Slave Node Info",
		FUNCTION_SEND_SLAVE_DATA:                   "Send Slave Data",
		FUNCTION_SET_SLAVE_LEARN_MODE:              "Set Slave Learn Mode",
		FUNCTION_GET_VIRTUAL_NODES:                 "Get Virtual Nodes",
		FUNCTION_IS_VIRTUAL_NODE:                   "Is Virtual Node",
		FUNCTION_SET_PROMISCUOUS_MODE:              "Set Promiscuous Mode",
	},
}

// Command class definitions from SDS12657-12-Z-Wave-Command-Class-Specification-A-M.pdf (3.4)
const (
	//COMMAND_CLASS_GROUPING_NAME [DEPRECATED]=0x7B // Grouping Name 1
	//COMMAND_CLASS_HAIL [DEPRECATED]=0x82 // Hail 1
	//COMMAND_CLASS_IP_CONFIGURATION [OBSOLETED]=0x9A // IP Configuratione 1
	//COMMAND_CLASS_PROPRIETARY  [DEPRECA TED]=0x88 // Proprietary 1
	//COMMAND_CLASS_REMOTE_ASSOCIATION_ACTIVATE  [OBSOLETED]=0x7C // Remote Association Activate
	COMMAND_CLASS_ALARM                             = 0x71 // Alarm 8
	COMMAND_CLASS_ANTITHEFT                         = 0x5D // Anti-theft 2
	COMMAND_CLASS_APPLICATION_CAPABILITY            = 0x57 // Application Capability1
	COMMAND_CLASS_APPLICATION_STATUS                = 0x22 // Application Status 1
	COMMAND_CLASS_ASSOCIATION                       = 0x85 // Association 2
	COMMAND_CLASS_ASSOCIATION_COMMAND_CONFIGURATION = 0x9B // Association Command Configuration 1
	COMMAND_CLASS_ASSOCIATION_GRP_INFO              = 0x59 // Association Group Information 3
	COMMAND_CLASS_BATTERY                           = 0x80 // Battery 1
	COMMAND_CLASS_CENTRAL_SCENE                     = 0x5B // Central Scene 3
	COMMAND_CLASS_CLOCK                             = 0x81 // Clock 1
	COMMAND_CLASS_CONFIGURATION                     = 0x70 // Configuration 4
	COMMAND_CLASS_CONTROLLER_REPLICATION            = 0x21 // Controller Replication 1
	COMMAND_CLASS_CRC_16_ENCAP                      = 0x56 // CRC-16 Encapsulation 1
	COMMAND_CLASS_DEVICE_RESET_LOCALLY              = 0x5A // Device Reset Locally 1
	COMMAND_CLASS_FIRMWARE_UPDATE_MD                = 0x7A // Firmware Update Meta Data 5
	COMMAND_CLASS_GEOGRAPHIC_LOCATION               = 0x8C // Geographic Location 1
	COMMAND_CLASS_INDICATOR                         = 0x87 // Indicator 2
	COMMAND_CLASS_IP_ASSOCIATION                    = 0x5C // IP Association 1
	COMMAND_CLASS_LANGUAGE                          = 0x89 // Language 1
	COMMAND_CLASS_MAILBOX                           = 0x69 // Mailbox 1
	COMMAND_CLASS_MANUFACTURER_PROPRIETARY          = 0x91 // Manufacturer Proprietary 1
	COMMAND_CLASS_MANUFACTURER_SPECIFIC             = 0x72 // Manufacturer Specific 2
	COMMAND_CLASS_MARK                              = 0xEF // Mark (Support/control mark) N/A
	COMMAND_CLASS_MULTI_CHANNEL                     = 0x60 // Multi Channel 4
	COMMAND_CLASS_MULTI_CHANNEL_ASSOCIATION         = 0x8E // Multi Channel Association 3
	COMMAND_CLASS_MULTI_COMMAND                     = 0x8F // Multi Command 1
	COMMAND_CLASS_NETWORK_MANAGEMENT_BASIC          = 0x4D // Network Management Basic Node 1
	COMMAND_CLASS_NETWORK_MANAGEMENT_INCLUSION      = 0x34 // Network Management Inclusion 1
	COMMAND_CLASS_NETWORK_MANAGEMENT_PRIMARY        = 0x54 // Network Management Primary 1
	COMMAND_CLASS_NETWORK_MANAGEMENT_PROXY          = 0x52 // Network Management Proxy  1
	COMMAND_CLASS_NODE_NAMING                       = 0x77 // Node Naming and Location 1
	COMMAND_CLASS_NON_INTEROPERABLE                 = 0xF0 // Non interoperable N/A
	COMMAND_CLASS_NOTIFICATION                      = 0x71 // Notification (former Alarm) 8
	COMMAND_CLASS_NO_OPERATION                      = 0x00 // No Operation 1

	// Configuration
	//	COMMAND_CLASS_SECURITY_SCHEME0_MARK = 0xF100 // Security Scheme 0 mark N/A
	COMMAND_CLASS_SCHEDULE          = 0x53 //	Schedule 3
	COMMAND_CLASS_SCREEN_ATTRIBUTES = 0x93 // Screen Attributes 2
	COMMAND_CLASS_SCREEN_MD         = 0x92 // Screen Meta Data 2
	COMMAND_CLASS_SECURITY          = 0x98 // Security 1
	COMMAND_CLASS_SUPERVISION       = 0x6C // Supervision 1
	COMMAND_CLASS_TIME              = 0x8A // Time 2
	COMMAND_CLASS_TIME_PARAMETERS   = 0x8B // Time Parameters 1
	COMMAND_CLASS_TRANSPORT_SERVICE = 0x55 // Transport Service 2
	COMMAND_CLASS_USER_CODE         = 0x63 // User Code 1
	COMMAND_CLASS_VERSION           = 0x86 // Version 2
	COMMAND_CLASS_WAKE_UP           = 0x84 // Wake Up 2
	COMMAND_CLASS_ZIP               = 0x23 // Z-Wave for IP 3
	COMMAND_CLASS_ZIP_6LOWPAN       = 0x4F // Z-Wave for IP, 6LoWPAN 8  (ITU-T G.9959 / IETF RFC 7428)
	COMMAND_CLASS_ZIP_GATEWAY       = 0x5F // Z-Wave for IP, Gateway 1
	COMMAND_CLASS_ZIP_NAMING        = 0x68 // Z-Wave for IP, Naming and Location 1
	COMMAND_CLASS_ZIP_ND            = 0x58 // Z-Wave for IP,-ND 1
	COMMAND_CLASS_ZIP_PORTAL        = 0x61 // Z-Wave for IP, Portal 1
	COMMAND_CLASS_ZWAVEPLUS_INFO    = 0x5E // Z-Wave Plus Info  2

	// Device specific
	//	COMMAND_CLASS_BASIC_WINDOW_COVERING            = 0x50 // Basic Window Covering 1  [OBSOLETED]
	//	COMMAND_CLASS_CLIMATE_CONTROL_SCHEDULE         = 0x46 // Climate Control Schedule 1 [DEPRECATED]
	//	COMMAND_CLASS_SENSOR_ALARM                     = 0x9C // Alarm Sensor 1  [DEPRECATED]
	//	COMMAND_CLASS_SENSOR_BINARY                    = 0x30 // Binary Sensor 2  [DEPRECATED]
	//	COMMAND_CLASS_SWITCH_TOGGLE_BINARY             = 0x28 // Binary Toggle Switch 1  [DEPRECATED]
	COMMAND_CLASS_BASIC                            = 0x20 // Basic 2
	COMMAND_CLASS_BASIC_TARIFF_INFO                = 0x36 // Basic Tariff Information1
	COMMAND_CLASS_DCP_CONFIG                       = 0x3A // DCP List Configuration1
	COMMAND_CLASS_DCP_MONITOR                      = 0x3B // DCP List Monitoring1
	COMMAND_CLASS_DOOR_LOCK                        = 0x62 // Door Lock3
	COMMAND_CLASS_DOOR_LOCK_LOGGING                = 0x4C // Door Lock Logging1
	COMMAND_CLASS_ENERGY_PRODUCTION                = 0x90 // Energy Production1
	COMMAND_CLASS_ENTRY_CONTROL                    = 0x6F // Entry Control1
	COMMAND_CLASS_HRV_CONTROL                      = 0x39 // HRV Control1
	COMMAND_CLASS_HRV_STATUS                       = 0x37 // HRV Status1
	COMMAND_CLASS_HUMIDITY_CONTROL_MODE            = 0x6D // Humidity Control Mode1
	COMMAND_CLASS_HUMIDITY_CONTROL_OPERATING_STATE = 0x6E // Humidity Control Operating State1
	COMMAND_CLASS_HUMIDITY_CONTROL_SETPOINT        = 0x64 // Humidity Control Setpoint1
	COMMAND_CLASS_IRRIGATION                       = 0x6B // Irrigation1
	COMMAND_CLASS_LOCK                             = 0x76 // Lock1
	COMMAND_CLASS_METER                            = 0x32 // Meter 4
	COMMAND_CLASS_METER_PULSE                      = 0x35 // Pulse Meter [DEPRECATED] 1
	COMMAND_CLASS_METER_TBL_CONFIG                 = 0x3C // Meter Table Configuration 1
	COMMAND_CLASS_METER_TBL_MONITOR                = 0x3D // Meter Table Monitor 2
	COMMAND_CLASS_METER_TBL_PUSH                   = 0x3E // Meter Table Push Configuration 1
	COMMAND_CLASS_MTP_WINDOW_COVERING              = 0x51 // Move To Position Window Covering [OBSOLETED] 1
	COMMAND_CLASS_POWERLEVEL                       = 0x73 // Powerlevel 1
	COMMAND_CLASS_PREPAYMENT                       = 0x3F // Prepayment 1
	COMMAND_CLASS_PREPAYMENT_ENCAPSULATION         = 0x41 // Prepayment Encapsulation 1
	COMMAND_CLASS_PROTECTION                       = 0x75 // Protection 2
	COMMAND_CLASS_RATE_TBL_CONFIG                  = 0x48 // Rate Table Configuration 1
	COMMAND_CLASS_RATE_TBL_MONITOR                 = 0x49 // Rate Table Monitoring 1
	COMMAND_CLASS_SCENE_ACTIVATION                 = 0x2B // Scene Activation 1
	COMMAND_CLASS_SCENE_ACTUATOR_CONF              = 0x2C // Scene Actuator Configuration 1
	COMMAND_CLASS_SCENE_CONTROLLER_CONF            = 0x2D // Scene Controller Configuration 1
	COMMAND_CLASS_SCHEDULE_ENTRY_LOCK              = 0x4E // Schedule Entry Lock [DEPRECATED] 3
	COMMAND_CLASS_SENSOR_CONFIGURATION             = 0x9E // Sensor Configuration [OBSOLETED] 1
	COMMAND_CLASS_SENSOR_MULTILEVEL                = 0x31 // Multilevel Sensor 10
	COMMAND_CLASS_SILENCE_ALARM                    = 0x9D // Alarm Silence 1
	COMMAND_CLASS_SIMPLE_AV_CONTROL                = 0x94 // Simple AV Control 4
	COMMAND_CLASS_SWITCH_ALL                       = 0x27 // All Switch 1
	COMMAND_CLASS_SWITCH_BINARY                    = 0x25 // Binary Switch 2
	COMMAND_CLASS_SWITCH_COLOR                     = 0x33 // Color Switch3
	COMMAND_CLASS_SWITCH_MULTILEVEL                = 0x26 // Multilevel Switch 4
	COMMAND_CLASS_SWITCH_TOGGLE_MULTILEVEL         = 0x29 // Multilevel Toggle Switch [DEPRECATED] 1
	COMMAND_CLASS_TARIFF_TBL_CONFIG                = 0x4A // Tariff Table Configuration 1
	COMMAND_CLASS_TARIFF_TBL_MONITOR               = 0x4B // Tariff Table Monitor 1
	COMMAND_CLASS_THERMOSTAT_FAN_MODE              = 0x44 // Thermostat Fan Mode 4
	COMMAND_CLASS_THERMOSTAT_FAN_STATE             = 0x45 // Thermostat Fan State 2
	COMMAND_CLASS_THERMOSTAT_MODE                  = 0x40 // Thermostat Mode 3
	COMMAND_CLASS_THERMOSTAT_OPERATING_STATE       = 0x42 // Thermostat Operating State 2
	COMMAND_CLASS_THERMOSTAT_SETBACK               = 0x47 // Thermostat Setback 1
	COMMAND_CLASS_THERMOSTAT_SETPOINT              = 0x43 // Thermostat Setpoint 3
	COMMAND_CLASS_WINDOW_COVERING                  = 0x6A // Window Covering 1
	COMMAND_CLASS_BARRIER_OPERATOR                 = 0x66 // Barrier Operator 1
)

var NamedCommandClass = &Named{
	"Command Class",
	map[byte]string{
		COMMAND_CLASS_ANTITHEFT:                         "Anti-theft",                        // 2
		COMMAND_CLASS_APPLICATION_CAPABILITY:            "Application Capabilit",             //y1
		COMMAND_CLASS_APPLICATION_STATUS:                "Application Status",                // 1
		COMMAND_CLASS_ASSOCIATION:                       "Association",                       // 2
		COMMAND_CLASS_ASSOCIATION_COMMAND_CONFIGURATION: "Association Command Configuration", // 1
		COMMAND_CLASS_ASSOCIATION_GRP_INFO:              "Association Group Information",     // 3
		COMMAND_CLASS_BATTERY:                           "Battery",                           // 1
		COMMAND_CLASS_CENTRAL_SCENE:                     "Central Scene",                     // 3
		COMMAND_CLASS_CLOCK:                             "Clock",                             // 1
		COMMAND_CLASS_CONFIGURATION:                     "Configuration",                     // 4
		COMMAND_CLASS_CONTROLLER_REPLICATION:            "Controller Replication",            // 1
		COMMAND_CLASS_CRC_16_ENCAP:                      "CRC-16 Encapsulation",              // 1
		COMMAND_CLASS_DEVICE_RESET_LOCALLY:              "Device Reset Locally",              // 1
		COMMAND_CLASS_FIRMWARE_UPDATE_MD:                "Firmware Update Meta Data",         // 5
		COMMAND_CLASS_GEOGRAPHIC_LOCATION:               "Geographic Location",               // 1
		COMMAND_CLASS_INDICATOR:                         "Indicator",                         // 2
		COMMAND_CLASS_IP_ASSOCIATION:                    "IP Association",                    // 1
		COMMAND_CLASS_LANGUAGE:                          "Language",                          // 1
		COMMAND_CLASS_MAILBOX:                           "Mailbox",                           // 1
		COMMAND_CLASS_MANUFACTURER_PROPRIETARY:          "Manufacturer Proprietary",          // 1
		COMMAND_CLASS_MANUFACTURER_SPECIFIC:             "Manufacturer Specific",             // 2
		COMMAND_CLASS_MARK:                              "Mark (Support/control mark) N",     // N/A
		COMMAND_CLASS_MULTI_CHANNEL:                     "Multi Channel",                     // 4
		COMMAND_CLASS_MULTI_CHANNEL_ASSOCIATION:         "Multi Channel Association",         // 3
		COMMAND_CLASS_MULTI_COMMAND:                     "Multi Command",                     // 1
		COMMAND_CLASS_NETWORK_MANAGEMENT_BASIC:          "Network Management Basic Node",     // 1
		COMMAND_CLASS_NETWORK_MANAGEMENT_INCLUSION:      "Network Management Inclusion",      // 1
		COMMAND_CLASS_NETWORK_MANAGEMENT_PRIMARY:        "Network Management Primary",        // 1
		COMMAND_CLASS_NETWORK_MANAGEMENT_PROXY:          "Network Management Proxy ",         // 1
		COMMAND_CLASS_NODE_NAMING:                       "Node Naming and Location",          // 1
		COMMAND_CLASS_NON_INTEROPERABLE:                 "Non interoperable N",               // N/A
		COMMAND_CLASS_NOTIFICATION:                      "Notification",                      // 8
		COMMAND_CLASS_NO_OPERATION:                      "No Operation",                      // 1

		// Configuration
		COMMAND_CLASS_SCHEDULE:          "Schedule",
		COMMAND_CLASS_SCREEN_ATTRIBUTES: "Screen",
		COMMAND_CLASS_SCREEN_MD:         "Screen",
		COMMAND_CLASS_SECURITY:          "Security",
		//	COMMAND_CLASS_SECURITY_SCHEME0_MARK: "Security", // will overflow since it's two bytes
		COMMAND_CLASS_SUPERVISION:       "Supervision",
		COMMAND_CLASS_TIME:              "Time",
		COMMAND_CLASS_TIME_PARAMETERS:   "Time",
		COMMAND_CLASS_TRANSPORT_SERVICE: "Transport",
		COMMAND_CLASS_USER_CODE:         "User",
		COMMAND_CLASS_VERSION:           "Version",
		COMMAND_CLASS_WAKE_UP:           "Wake",
		COMMAND_CLASS_ZIP:               "Z-Wave For IP",
		COMMAND_CLASS_ZIP_NAMING:        "Z-Wave For IP, Naming and Location",
		COMMAND_CLASS_ZIP_ND:            "Z-Wave For IP,-ND",
		COMMAND_CLASS_ZIP_6LOWPAN:       "Z-Wave For IP, 6LoWPAN (ITU-T G.9959 / IETF RFC 7428)",
		COMMAND_CLASS_ZIP_GATEWAY:       "Z-Wave For IP, Gateway",
		COMMAND_CLASS_ZIP_PORTAL:        "Z-Wave For IP, Portal",
		COMMAND_CLASS_ZWAVEPLUS_INFO:    "Z-Wave Plus Info",

		// Device specific
		COMMAND_CLASS_BASIC:                            "Basic",
		COMMAND_CLASS_BASIC_TARIFF_INFO:                "Basic Tariff Information",
		COMMAND_CLASS_DCP_CONFIG:                       "DCP List Configuration",
		COMMAND_CLASS_DCP_MONITOR:                      "DCP List Monitoring",
		COMMAND_CLASS_DOOR_LOCK:                        "Door Lock",
		COMMAND_CLASS_DOOR_LOCK_LOGGING:                "Door Lock Logging",
		COMMAND_CLASS_ENERGY_PRODUCTION:                "Energy Production",
		COMMAND_CLASS_ENTRY_CONTROL:                    "Entry Control",
		COMMAND_CLASS_HRV_CONTROL:                      "HRV Control",
		COMMAND_CLASS_HRV_STATUS:                       "HRV Status",
		COMMAND_CLASS_HUMIDITY_CONTROL_MODE:            "Humidity Control Mode",
		COMMAND_CLASS_HUMIDITY_CONTROL_OPERATING_STATE: "Humidity Control Operating State",
		COMMAND_CLASS_HUMIDITY_CONTROL_SETPOINT:        "Humidity Control Setpoint",
		COMMAND_CLASS_IRRIGATION:                       "Irrigation",
		COMMAND_CLASS_LOCK:                             "Lock",
		COMMAND_CLASS_METER:                            "Meter",
		COMMAND_CLASS_METER_PULSE:                      "Pulse Meter [DEPRECATED]",
		COMMAND_CLASS_METER_TBL_CONFIG:                 "Meter Table Configuration",
		COMMAND_CLASS_METER_TBL_MONITOR:                "Meter Table Monitor",
		COMMAND_CLASS_METER_TBL_PUSH:                   "Meter Table Push Configuration",
		COMMAND_CLASS_MTP_WINDOW_COVERING:              "Move To Position Window Covering [OBSOLETED]",
		COMMAND_CLASS_POWERLEVEL:                       "Powerlevel",
		COMMAND_CLASS_PREPAYMENT:                       "Prepayment",
		COMMAND_CLASS_PREPAYMENT_ENCAPSULATION:         "Prepayment Encapsulation",
		COMMAND_CLASS_PROTECTION:                       "Protection",
		COMMAND_CLASS_RATE_TBL_CONFIG:                  "Rate Table Configuration",
		COMMAND_CLASS_RATE_TBL_MONITOR:                 "Rate Table Monitoring",
		COMMAND_CLASS_SCENE_ACTIVATION:                 "Scene Activation",
		COMMAND_CLASS_SCENE_ACTUATOR_CONF:              "Scene Actuator Configuration",
		COMMAND_CLASS_SCENE_CONTROLLER_CONF:            "Scene Controller Configuration",
		COMMAND_CLASS_SCHEDULE_ENTRY_LOCK:              "Schedule Entry Lock [DEPRECATED]",
		COMMAND_CLASS_SENSOR_CONFIGURATION:             "Sensor Configuration [OBSOLETED]",
		COMMAND_CLASS_SENSOR_MULTILEVEL:                "Multilevel Sensor ",
		COMMAND_CLASS_SILENCE_ALARM:                    "Alarm Silence",
		COMMAND_CLASS_SIMPLE_AV_CONTROL:                "Simple AV Control",
		COMMAND_CLASS_SWITCH_ALL:                       "All Switch",
		COMMAND_CLASS_SWITCH_BINARY:                    "Binary Switch",
		COMMAND_CLASS_SWITCH_COLOR:                     "Color Switch",
		COMMAND_CLASS_SWITCH_MULTILEVEL:                "Multilevel Switch",
		COMMAND_CLASS_SWITCH_TOGGLE_MULTILEVEL:         "Multilevel Toggle Switch [DEPRECATED]",
		COMMAND_CLASS_TARIFF_TBL_CONFIG:                "Tariff Table Configuration",
		COMMAND_CLASS_TARIFF_TBL_MONITOR:               "Tariff Table Monitor",
		COMMAND_CLASS_THERMOSTAT_FAN_MODE:              "Thermostat Fan Mode",
		COMMAND_CLASS_THERMOSTAT_FAN_STATE:             "Thermostat Fan State",
		COMMAND_CLASS_THERMOSTAT_MODE:                  "Thermostat Mode",
		COMMAND_CLASS_THERMOSTAT_OPERATING_STATE:       "Thermostat Operating State",
		COMMAND_CLASS_THERMOSTAT_SETBACK:               "Thermostat Setback",
		COMMAND_CLASS_THERMOSTAT_SETPOINT:              "Thermostat Setpoint",
		COMMAND_CLASS_WINDOW_COVERING:                  "Window Covering",
		COMMAND_CLASS_BARRIER_OPERATOR:                 "Barrier Operator",
	},
}

const (
	NOTIFICATION_GET              = 0x04 // NOTIFICATION is ALARM in the spec
	NOTIFICATION_REPORT           = 0x05
	NOTIFICATION_SUPPORTED_GET    = 0x07
	NOTIFICATION_SUPPORTED_REPORT = 0x08
)

var NamedNotificationCommand = &Named{
	"Notification Command",
	map[byte]string{
		NOTIFICATION_GET:              "Get",
		NOTIFICATION_REPORT:           "Report",
		NOTIFICATION_SUPPORTED_GET:    "Supported Get",
		NOTIFICATION_SUPPORTED_REPORT: "Supported Report",
	},
}

// 4.3.4 Alarm Report Command, Z-Wave Alarm Type
const (
	ZALARM_TYPE_RESERVED         = 0x00
	ZALARM_TYPE_SMOKE            = 0x01
	ZALARM_TYPE_CARBONMONOXIDE   = 0x02
	ZALARM_TYPE_CARBONDIOXIDE    = 0x03
	ZALARM_TYPE_HEAT             = 0x04
	ZALARM_TYPE_WATER            = 0x05
	ZALARM_TYPE_ACCESS_CONTROL   = 0x06
	ZALARM_TYPE_BURGLAR          = 0x07
	ZALARM_TYPE_POWER_MANAGEMENT = 0x08
	ZALARM_TYPE_SYSTEM           = 0x09
	ZALARM_TYPE_EMERGENCY        = 0x0a
	ZALARM_TYPE_CLOCK            = 0x0b
	// 0x0c -> 0xfe Reserved
	ZALARM_TYPE_APPLIANCE    = 0x0c // from openwave project
	ZALARM_TYPE_HOMEHEALTH   = 0x0d
	ZALARM_TYPE_COUNT        = 0x0e
	ZALARM_TYPE_RETURN_FIRST = 0xff
)

var NamedZAlarmType = &Named{
	"Z-Wave Alarm Type",
	map[byte]string{
		ZALARM_TYPE_RESERVED:         "Reserved",
		ZALARM_TYPE_SMOKE:            "Smoke",
		ZALARM_TYPE_CARBONMONOXIDE:   "Carbonmonoxide",
		ZALARM_TYPE_CARBONDIOXIDE:    "Carbondioxide",
		ZALARM_TYPE_HEAT:             "Heat",
		ZALARM_TYPE_WATER:            "Water",
		ZALARM_TYPE_ACCESS_CONTROL:   "Access Control",
		ZALARM_TYPE_BURGLAR:          "Burglar",
		ZALARM_TYPE_POWER_MANAGEMENT: "Power Management",
		ZALARM_TYPE_SYSTEM:           "System",
		ZALARM_TYPE_EMERGENCY:        "Emergency",
		ZALARM_TYPE_CLOCK:            "Clock",
		ZALARM_TYPE_APPLIANCE:        "Appliance",
		ZALARM_TYPE_HOMEHEALTH:       "Homehealth",
		ZALARM_TYPE_COUNT:            "Count",
		ZALARM_TYPE_RETURN_FIRST:     "Return First",
	},
}

// Z-Wave Alarm Event (8 bits) and Event Parameters (N Byte)
const (
	BURGLAR_RESERVED                      = 0x00
	BURGLAR_INTRUSION                     = 0x01
	BURGLAR_INTRUSION_UNKNOWN_LOCATION    = 0x02
	BURGLAR_TAMPERING_COVERING_REMOVED    = 0x03
	BURGLAR_TAMPERING_INVALID             = 0x04
	BURGLAR_GLASS_BROKEN                  = 0x05
	BURGLAR_GLASS_BROKEN_UNKNOWN_LOCATION = 0x06
	BURGLAR_UNKNOWN_EVENT                 = 0x0f
	BURGLAR_RESERVED_LAST                 = 0xff
	// 0x07-0xfd Reserved
	BURGLAR_TAMPERING = 0x07 // when eg. shaking the fibaro motion sensor
	BURGLAR_FIBARO    = 0x08 // comes from time to tome from fibaro motion sensor
)

var NamedBurglarEvent = &Named{
	"Burglar Event",
	map[byte]string{
		BURGLAR_RESERVED:                      "Reserved",
		BURGLAR_INTRUSION:                     "Intrusion",
		BURGLAR_INTRUSION_UNKNOWN_LOCATION:    "Intrusion Unknown Location",
		BURGLAR_TAMPERING_COVERING_REMOVED:    "Tampering Covering Removed",
		BURGLAR_TAMPERING_INVALID:             "Tampering Invalid",
		BURGLAR_GLASS_BROKEN:                  "Glass Broken",
		BURGLAR_GLASS_BROKEN_UNKNOWN_LOCATION: "Glass Broken Unknown Location",
		BURGLAR_UNKNOWN_EVENT:                 "Unknown Event",
		BURGLAR_RESERVED_LAST:                 "Reserved",
		// Other reserved
		BURGLAR_TAMPERING: "Tampering",
		BURGLAR_FIBARO:    "Fibaro motion sensor something",
	},
}

const (
	ZALARM_STATUS_OFF = 0x00
	ZALARM_STATUS_ON  = 0xff
)

var NamedZAlarmStatus = &Named{
	"Z-Wave Alarm Status",
	map[byte]string{
		ZALARM_STATUS_ON:  "On",
		ZALARM_STATUS_OFF: "Off",
	},
}
