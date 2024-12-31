package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json_input := []byte(`{
  "endpoints": [
    {
      "hostname": "preimer1",
      "display_name": "",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "a0:48:1c:9b:1b:54",
          "ips": [
            "172.25.18.64"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "",
      "xdr_device_id": "d6590e25-3929-4edb-b952-a6b3fa9070c6",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86",
        "version": "10.0 (Build 19042)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1731508334,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": "1.2.0.5608"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "ncie:recording:windows:module_error:0xd0000001"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1731087495,
      "xes_ei_update_time": 1731507796,
      "epp_tracking_time": 0,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 47,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-03",
              "name": "Extended non-use"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-04",
              "name": "Very few"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-08",
              "name": "Entertainment"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "LTMABJO013",
      "display_name": "LTMABJO013",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "b86d595e-b7f0-4b59-a6db-eb0475067fa8",
        "ws_group_id": "",
        "name": "Office-computers",
        "type": "Folder",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "c1f3b6ae-1cd4-461a-9e55-d0aa910fa9f0",
        "name": "ARM_POL_COMPUTERS_ALL_MAIN"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "20:79:18:b1:90:e0",
          "ips": [
            "10.212.25.30"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.212.25.30",
      "xdr_device_id": "844ed3be-9ed0-45fa-91e1-b4018a230973",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19045)",
        "kernel_version": ""
      },
      "cpu_arch": "64-bit",
      "last_logged_on_user": "ARMOR\\zhadi",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1733910494,
      "epp_component_module_status": "off",
      "epp_component_module_last_connected_time": 1733908193,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "14.0.14203",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "536870944",
        "1208221953",
        "1208222100",
        "1208221798",
        "1208222211",
        "1208222245",
        "1208222096",
        "1090519040",
        "1208222021",
        "1073741840",
        "1208222099",
        "1073741856",
        "1208221808",
        "512",
        "1208221779",
        "1208221988",
        "536870976",
        "1208221844",
        "1208222229",
        "1208090624",
        "1208221826"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:tool_error:106"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975321,
      "xes_ei_update_time": 1733830588,
      "epp_tracking_time": 1733830602,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "outdated",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-03",
              "name": "Extended non-use"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-04",
              "name": "Very few"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVUSAMOPWINFRA.armor.fr",
      "display_name": "SVUSAMOPWINFRA",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "55d975f7-8594-b591-8c2a-9b816f10c38b",
        "ws_group_id": "1777",
        "name": "INFRA",
        "type": "vcenter-folder-network",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Network adapter 1",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:0C:29:25:A0:09",
          "ips": []
        },
        {
          "name": "VL_O_SRV_PRD1",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:0C:29:25:A0:09",
          "ips": [
            "172.25.1.3"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.25.1.3",
      "xdr_device_id": "6c9dc9de-8078-4569-1b57-f8d0f55c2b26",
      "ws_device_id": "4773",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1727250683,
      "xes_ei_update_time": 1734071039,
      "epp_tracking_time": 1698328556,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "pfrwaxsql",
      "display_name": "",
      "description": "",
      "type": "server",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:50:56:a3:1b:bf",
          "ips": [
            "172.27.14.5"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "[\"172.27.14.5\"]",
      "xdr_device_id": "7a52e134-93af-4aee-95ea-4ef1b22bbade",
      "ws_device_id": "",
      "os": {
        "name": "Windows Server 2008 R2",
        "platform": "windows",
        "arch": "x86_64",
        "version": "6.1 (Build 7601)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": "1.2.0.4648"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "c1_0x1006"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975323,
      "xes_ei_update_time": 1733449707,
      "epp_tracking_time": 0,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 63,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "pcornwrds01.armor.fr",
      "display_name": "pcornwrds01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "8f54200e-0941-84b4-9379-10cfd1012248",
        "ws_group_id": "1523",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:7D:CC",
          "ips": [
            "10.57.1.11"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.57.1.11",
      "xdr_device_id": "3690fe2f-d764-5b1d-b76c-092a42f403f9",
      "ws_device_id": "3767",
      "os": {
        "name": "Microsoft Windows Server 2016",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 14393)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1733336447,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975309,
      "xes_ei_update_time": 1733333377,
      "epp_tracking_time": 1698320404,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "outdated",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "pmxsqwaps",
      "display_name": "",
      "description": "",
      "type": "server",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:50:56:9e:21:7f",
          "ips": [
            "10.52.1.7"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "",
      "xdr_device_id": "356d860d-8d75-4da5-8816-52bba3d2752a",
      "ws_device_id": "",
      "os": {
        "name": "Windows Server 2012 R2",
        "platform": "windows",
        "arch": "x86_64",
        "version": "6.3 (Build 9600)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "off",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "2",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1208222296"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975323,
      "xes_ei_update_time": 1733816607,
      "epp_tracking_time": 1698330186,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 79,
      "risk_level": "high",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PMABJWOPTITIME.armor.fr",
      "display_name": "PMABJWOPTITIME",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "a9413f60-998c-57d8-6170-e647caac4994",
        "ws_group_id": "1619",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:63:B1",
          "ips": [
            "10.212.1.7"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.212.1.7",
      "xdr_device_id": "eced764b-3bc4-4939-858a-20e5daaf28d9",
      "ws_device_id": "3980",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1727250683,
      "xes_ei_update_time": 1733997918,
      "epp_tracking_time": 1698319967,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 56,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "dc-pl1nodea.iimak.local",
      "display_name": "DC-PL1NODEA",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "fe48a939-1875-cecf-0fb7-37cd7f2308ea",
        "ws_group_id": "0",
        "name": "Computers",
        "type": "root",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "102",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:26:b9:74:f9:38",
          "ips": [
            "10.1.12.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.1.12.1",
      "xdr_device_id": "e4453674-3658-4083-b62f-6fd5eb8537dc",
      "ws_device_id": "4507",
      "os": {
        "name": "Microsoft Windows 7",
        "platform": "windows",
        "arch": "x86_64",
        "version": "6.1 (Build 7601)",
        "kernel_version": ""
      },
      "cpu_arch": "64-bit",
      "last_logged_on_user": "INTRANET_DOMAIN\\data1",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734051958,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.0.6313",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "not-supported"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "not-supported"
        }
      ],
      "installed_components": [
        "1082130432",
        "536870944",
        "1208221953",
        "1208222100",
        "1208221798",
        "1208222211",
        "1208222245",
        "1208222096",
        "1090519040",
        "1208222021",
        "1073741840",
        "1208222099",
        "1073741856",
        "1208221808",
        "512",
        "1208221779",
        "1208221988",
        "536870976",
        "1208221844",
        "1208222229",
        "1208090624",
        "1208221826"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_precheck:windows:tool_error:255"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_precheck",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1734057343,
      "xes_ei_update_time": 1732693270,
      "epp_tracking_time": 1734060800,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "outdated",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 66,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "WCOFS",
      "display_name": "WCOFS",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "1a92db2b-7cce-4d50-9d40-872577090b82",
        "ws_group_id": "",
        "name": "Intranet_domain",
        "type": "Folder",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "50:65:f3:49:22:55",
          "ips": [
            "192.168.13.8"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "192.168.13.8",
      "xdr_device_id": "ff2ec6da-9390-4328-a4ec-736b980664f2",
      "ws_device_id": "",
      "os": {
        "name": "Windows 7",
        "platform": "windows",
        "arch": "x86",
        "version": "6.1 (Build 7601)",
        "kernel_version": ""
      },
      "cpu_arch": "32-bit",
      "last_logged_on_user": "INTRANET_DOMAIN\\wcclip",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "off",
      "epp_component_module_last_connected_time": 1734037607,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "14.0.12032",
        "endpoint_sensor": ""
      },
      "last_scan_time": 1733518407,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "536870944",
        "1208222100",
        "1208221798",
        "1208222211",
        "1208221731",
        "1208222096",
        "1090519040",
        "1208222021",
        "1073741840",
        "1208222099",
        "1073741856",
        "1208221808",
        "512",
        "1208221846",
        "1208221779",
        "1208221801",
        "1208221988",
        "536870976",
        "1208221800",
        "1208090624"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_precheck:windows:xbc_error:10001"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_precheck",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975310,
      "xes_ei_update_time": 1732637842,
      "epp_tracking_time": 1720415956,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "outdated",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "pushewsec01",
      "display_name": "",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "f8:b1:56:b9:78:fd",
          "ips": [
            "10.1.15.2"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "[\"10.1.15.2\"]",
      "xdr_device_id": "59ebe8c9-ce26-4e5b-8e54-d8358230e665",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19044)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1208222296"
      ],
      "recommended_actions": {
        "errors": [
          "xbc_error_10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975322,
      "xes_ei_update_time": 1733860920,
      "epp_tracking_time": 0,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 72,
      "risk_level": "high",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-08",
              "name": "Entertainment"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "JCARDVM",
      "display_name": "JCARDVM",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "1a92db2b-7cce-4d50-9d40-872577090b82",
        "ws_group_id": "",
        "name": "Intranet_domain",
        "type": "Folder",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "c1f3b6ae-1cd4-461a-9e55-d0aa910fa9f0",
        "name": "ARM_POL_COMPUTERS_ALL_MAIN"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:50:56:91:7e:ae",
          "ips": [
            "172.25.1.97"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.25.1.97",
      "xdr_device_id": "3041c50b-245d-4ce1-84fe-3035dbb49c9d",
      "ws_device_id": "",
      "os": {
        "name": "Windows 7",
        "platform": "windows",
        "arch": "x86",
        "version": "6.1 (Build 7601)",
        "kernel_version": ""
      },
      "cpu_arch": "32-bit",
      "last_logged_on_user": "INTRANET_DOMAIN\\joel.hopkins",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734045098,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734051162,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "14.0.10349",
        "endpoint_sensor": ""
      },
      "last_scan_time": 1733506732,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208222100",
        "536870944",
        "1208222211",
        "1208221798",
        "1208221731",
        "1208222096",
        "1090519040",
        "1208222021",
        "1073741840",
        "1208222099",
        "1073741856",
        "1208221808",
        "512",
        "1208221846",
        "1208221779",
        "1208221801",
        "1208221988",
        "536870976",
        "1208221800",
        "1208222229",
        "1208090624"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_precheck:windows:xbc_error:10001"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_precheck",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975307,
      "xes_ei_update_time": 1732692895,
      "epp_tracking_time": 1728572723,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 44,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "pushewaps.armor.fr",
      "display_name": "pushewaps",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "87f5c936-682e-58b1-7eb8-bc5928f7310a",
        "ws_group_id": "1630",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "10",
        "name": "ARM_POL_SRV_PRO_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:9E:35:B9",
          "ips": [
            "10.1.1.5"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.1.1.5",
      "xdr_device_id": "53a73380-1cf7-448c-8901-cc7b9bcc34ff",
      "ws_device_id": "3998",
      "os": {
        "name": "Microsoft Windows Server 2012 R2",
        "platform": "windows",
        "arch": "x86_64",
        "version": "6.3 (Build 9600)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975306,
      "xes_ei_update_time": 1733812592,
      "epp_tracking_time": 1698319967,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 55,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PSGSGWAPS.armor.fr",
      "display_name": "PSGSGWAPS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "b640ddf5-30c8-8d0b-9d04-7dbe9495e222",
        "ws_group_id": "1642",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "10",
        "name": "ARM_POL_SRV_PRO_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:24:E8",
          "ips": [
            "10.65.1.21"
          ]
        }
      ],
      "vdi": {
        "type": ""
      },
      "last_ip_used": "10.65.1.21",
      "xdr_device_id": "24fb0e06-0896-4cae-95e1-7b40b38392b8",
      "ws_device_id": "4463",
      "os": {
        "name": "Microsoft Windows Server 2012 R2",
        "platform": "windows",
        "arch": "x86_64",
        "version": "6.3 (Build 9600)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "unknown",
      "last_connected_time": 0,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": [
          "tip-sensor-update-required-download-and-install-agent-server-workload-protection"
        ]
      },
      "xes_deployment_status": "",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1710111637,
      "xes_ei_update_time": 0,
      "epp_tracking_time": 1698324357,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 57,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PCNZHWT1HR.armor.fr",
      "display_name": "PCNZHWT1HR",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "976a44c3-9055-19b3-c8d9-47550e7f8b70",
        "ws_group_id": "1528",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "101",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY_ACTIVITY-MON-DISABLED"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "1-ARMOR",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:8B:41:0D",
          "ips": [
            "10.85.1.19"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.85.1.19",
      "xdr_device_id": "2c730187-1fc8-4788-8478-f6776c45bd4d",
      "ws_device_id": "4252",
      "os": {
        "name": "Microsoft Windows Server 2016",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 14393)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989976,
      "xes_ei_update_time": 1733833730,
      "epp_tracking_time": 1698324405,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 59,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVCNSHOPWINFRA.armor.fr",
      "display_name": "SVCNSHOPWINFRA",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "3b1527d8-535b-b22e-ae02-8c311cfd6330",
        "ws_group_id": "1651",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "6304c6e24421b4519ac3221c",
        "xes_group_name": "all endpoints"
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:82:16",
          "ips": [
            "fe80::672:300b:50c1:f1a6%8",
            "10.87.1.3"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.87.1.3",
      "xdr_device_id": "457df4d7-d7ee-62ca-6dcb-c54cb616b002",
      "ws_device_id": "4049",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734050106,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989971,
      "xes_ei_update_time": 1733445505,
      "epp_tracking_time": 1698319967,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVUSHEOPWINFRA.armor.fr",
      "display_name": "SVUSHEOPWINFRA",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "152f77f2-0398-207f-5674-889bcdac89d8",
        "ws_group_id": "1627",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:C7:8C",
          "ips": [
            "10.1.1.3",
            "fe80::fb11:5fac:a6cb:9119%8"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.1.1.3",
      "xdr_device_id": "01132082-a7c5-e647-01d7-123c002d48af",
      "ws_device_id": "4439",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1727250683,
      "xes_ei_update_time": 1733988447,
      "epp_tracking_time": 1698324707,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 59,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-03",
              "name": "Several regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVSGSGOPWINFRA.armor.fr",
      "display_name": "SVSGSGOPWINFRA",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "97780b6e-21ba-94fd-fe91-248cface5eea",
        "ws_group_id": "1639",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet1",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:7F:3C",
          "ips": [
            "fe80::e923:cd81:a5d:3460%15",
            "10.65.250.125"
          ]
        },
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:2D:76",
          "ips": [
            "10.65.1.3"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.65.1.3",
      "xdr_device_id": "3a7c0109-eb02-4682-9757-7162ac36de29",
      "ws_device_id": "4177",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989806,
      "xes_ei_update_time": 1733989507,
      "epp_tracking_time": 1698319967,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 59,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-01",
              "name": "Design and development"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRBFOPWINFRA.armor.fr",
      "display_name": "SVFRBFOPWINFRA",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "32aeadb0-3870-37c6-cfba-fb1fed532ae8",
        "ws_group_id": "1949",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Network adapter 1",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:57:B6",
          "ips": []
        },
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:57:B6",
          "ips": [
            "172.19.101.50",
            "172.19.1.5",
            "fe80::6414:82e9:cf63:d0d7%8",
            "172.19.1.3"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.19.1.3",
      "xdr_device_id": "a842a90e-827e-4008-ad1b-9a05bdf7ce61",
      "ws_device_id": "4946",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1727250683,
      "xes_ei_update_time": 1733983577,
      "epp_tracking_time": 1712753807,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 59,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-03",
              "name": "Cloud services"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-03",
              "name": "Several regular"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "TBFRCH029",
      "display_name": "TBFRCH029",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "b93f8d3a-026f-4974-90bc-6b978d6af2be",
        "ws_group_id": "",
        "name": "Workgroup",
        "type": "Native",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "c1f3b6ae-1cd4-461a-9e55-d0aa910fa9f0",
        "name": "ARM_POL_COMPUTERS_ALL_MAIN"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "e4:fd:45:5f:24:63",
          "ips": [
            "172.30.131.164"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.30.131.164",
      "xdr_device_id": "bcd2b3ae-91e9-4d7e-be38-5350f119c599",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19045)",
        "kernel_version": ""
      },
      "cpu_arch": "64-bit",
      "last_logged_on_user": "ARMOR\\Jade",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734062074,
      "epp_component_module_status": "off",
      "epp_component_module_last_connected_time": 1734062090,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "14.0.14203",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 1733486771,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208222100",
        "1208221953",
        "1208221798",
        "1208222245",
        "1208222021",
        "1090519040",
        "1073741840",
        "1073741856",
        "512",
        "1208221779",
        "1208222229",
        "1208090624",
        "536870944",
        "1208222211",
        "1208222096",
        "1208222296",
        "1208222099",
        "1208221808",
        "1208221988",
        "536870976",
        "1208221844",
        "1208221826"
      ],
      "recommended_actions": {
        "errors": [
          "c1_0x1101"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975311,
      "xes_ei_update_time": 1733844576,
      "epp_tracking_time": 1700141539,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 53,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PCNZHWAD01.armor.fr",
      "display_name": "PCNZHWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "f69a4b9b-37d7-517d-483c-ad686a36230e",
        "ws_group_id": "1526",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:FA:C9",
          "ips": [
            "10.85.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.85.1.1",
      "xdr_device_id": "6b8d7c9b-e691-ce61-57e9-2d4cb528cb26",
      "ws_device_id": "3779",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734069338,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733990051,
      "xes_ei_update_time": 1727976725,
      "epp_tracking_time": 1698320485,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "IC",
          "name": "IP segment connections",
          "tags": [
            {
              "id": "IC-02",
              "name": "Few cross-segment"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-12",
              "name": "LDAP"
            },
            {
              "id": "DR-14",
              "name": "KERBEROS"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "DFRDCWPRNSRV",
      "display_name": "",
      "description": "",
      "type": "server",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:50:56:87:27:61",
          "ips": [
            "172.27.251.34"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "[\"172.27.251.36\"]",
      "xdr_device_id": "8fda7708-2d21-464c-b882-c318238198e3",
      "ws_device_id": "",
      "os": {
        "name": "Windows Server, version 1809",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": "1.2.0.4757"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975323,
      "xes_ei_update_time": 1733865271,
      "epp_tracking_time": 0,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-01",
              "name": "Design and development"
            },
            {
              "id": "AU-05",
              "name": "Finance"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PSGSGWAD01.armor.fr",
      "display_name": "PSGSGWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "97780b6e-21ba-94fd-fe91-248cface5eea",
        "ws_group_id": "1639",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:9C:47",
          "ips": [
            "10.65.1.1"
          ]
        }
      ],
      "vdi": {
        "type": ""
      },
      "last_ip_used": "10.65.1.1",
      "xdr_device_id": "14cec51c-30ae-5f31-5c9a-5a32699c5b0e",
      "ws_device_id": "4460",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "unknown",
      "last_connected_time": 0,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": [
          "tip-sensor-update-required-download-and-install-agent-server-workload-protection"
        ]
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989891,
      "xes_ei_update_time": 0,
      "epp_tracking_time": 1698319967,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "IC",
          "name": "IP segment connections",
          "tags": [
            {
              "id": "IC-02",
              "name": "Few cross-segment"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-12",
              "name": "LDAP"
            },
            {
              "id": "DR-14",
              "name": "KERBEROS"
            },
            {
              "id": "DR-UNK",
              "name": "Unknown"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PTRISWAD01.armor.fr",
      "display_name": "PTRISWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "8fd21971-6995-b1df-5716-ab413452a432",
        "ws_group_id": "1645",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:78:78",
          "ips": [
            "10.90.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.90.1.1",
      "xdr_device_id": "c85a7add-1ab1-a7b1-82c2-ea7a61ddb678",
      "ws_device_id": "4468",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734051223,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733990082,
      "xes_ei_update_time": 1724902531,
      "epp_tracking_time": 1698319967,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "IC",
          "name": "IP segment connections",
          "tags": [
            {
              "id": "IC-02",
              "name": "Few cross-segment"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-14",
              "name": "KERBEROS"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PCORNWAD01",
      "display_name": "",
      "description": "",
      "type": "server",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:50:56:93:54:75",
          "ips": [
            "10.57.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "[\"10.57.1.1\"]",
      "xdr_device_id": "cd4c8f92-b50c-4327-a702-d72639fdea75",
      "ws_device_id": "",
      "os": {
        "name": "Windows Server, version 1809",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734043578,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975323,
      "xes_ei_update_time": 1724902531,
      "epp_tracking_time": 0,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "unknown",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-03",
              "name": "Extended non-use"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PMXSQWAD01.armor.fr",
      "display_name": "PMXSQWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "b57dc183-1609-1839-0545-05c50415a4cc",
        "ws_group_id": "1612",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "101",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY_ACTIVITY-MON-DISABLED"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:BF:7E",
          "ips": [
            "10.52.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.52.1.1",
      "xdr_device_id": "2a608a43-c784-4a48-a3dc-dda41da30243",
      "ws_device_id": "4408",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734050309,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975312,
      "xes_ei_update_time": 1724902533,
      "epp_tracking_time": 1698319967,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "IC",
          "name": "IP segment connections",
          "tags": [
            {
              "id": "IC-02",
              "name": "Few cross-segment"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-12",
              "name": "LDAP"
            },
            {
              "id": "DR-14",
              "name": "KERBEROS"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "pcnzhwfiles.armor.fr",
      "display_name": "pcnzhwfiles",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "976a44c3-9055-19b3-c8d9-47550e7f8b70",
        "ws_group_id": "1528",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "101",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY_ACTIVITY-MON-DISABLED"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:8B:68:15",
          "ips": [
            "10.85.1.5"
          ]
        }
      ],
      "vdi": {
        "type": ""
      },
      "last_ip_used": "10.85.1.5",
      "xdr_device_id": "1a4bd006-1a68-2868-8d04-043472c8c6cf",
      "ws_device_id": "4251",
      "os": {
        "name": "Microsoft Windows Server 2016",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 14393)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "unknown",
      "last_connected_time": 0,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_one_cpu_core_110"
        ],
        "tips": [
          "tip-sensor-disabled-download-and-install-agent-server-workload-protection"
        ]
      },
      "xes_deployment_status": "",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1705532620,
      "xes_ei_update_time": 0,
      "epp_tracking_time": 1698330168,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PKENAWAD01.armor.fr",
      "display_name": "PKENAWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "430fa561-8701-dde9-d07a-6675f002f895",
        "ws_group_id": "1622",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "101",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY_ACTIVITY-MON-DISABLED"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:E7:85",
          "ips": [
            "10.25.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.25.1.1",
      "xdr_device_id": "350c9765-712c-4851-a8a6-aa20569db057",
      "ws_device_id": "4426",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734059545,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975312,
      "xes_ei_update_time": 1724902528,
      "epp_tracking_time": 1698328624,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 59,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "IC",
          "name": "IP segment connections",
          "tags": [
            {
              "id": "IC-02",
              "name": "Few cross-segment"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-14",
              "name": "KERBEROS"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PMABJWAD01.armor.fr",
      "display_name": "PMABJWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "80e66f93-8f0a-c1f6-03d2-adf174cffe27",
        "ws_group_id": "1617",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:4E:22",
          "ips": [
            "10.212.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.212.1.1",
      "xdr_device_id": "c7f34ea3-7e14-4063-ab1a-2a98f23cde41",
      "ws_device_id": "3977",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734060465,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733987773,
      "xes_ei_update_time": 1724902534,
      "epp_tracking_time": 1698319967,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-12",
              "name": "LDAP"
            },
            {
              "id": "DR-14",
              "name": "KERBEROS"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PCATOWAD01",
      "display_name": "",
      "description": "",
      "type": "server",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "00:50:56:93:66:cd",
          "ips": [
            "10.11.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "[\"10.11.1.1\"]",
      "xdr_device_id": "68a5c220-cc30-4eb4-9093-540a8e01329e",
      "ws_device_id": "",
      "os": {
        "name": "Windows Server, version 1809",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734068753,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "inspect_error_connect_to_server_failed_146"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975323,
      "xes_ei_update_time": 1724902530,
      "epp_tracking_time": 0,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "unknown",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-03",
              "name": "Extended non-use"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-01",
              "name": "One"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "TBFRCH031",
      "display_name": "TBFRCH031",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "b93f8d3a-026f-4974-90bc-6b978d6af2be",
        "ws_group_id": "",
        "name": "Workgroup",
        "type": "Native",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "c1f3b6ae-1cd4-461a-9e55-d0aa910fa9f0",
        "name": "ARM_POL_COMPUTERS_ALL_MAIN"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "30:05:05:d8:26:f1",
          "ips": [
            "172.30.130.97"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.30.130.97",
      "xdr_device_id": "1930ad6d-21c0-487d-83c8-cf6d7053d980",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19045)",
        "kernel_version": ""
      },
      "cpu_arch": "64-bit",
      "last_logged_on_user": "ARMOR\\fzongo",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1733999115,
      "epp_component_module_status": "off",
      "epp_component_module_last_connected_time": 1733999130,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "14.0.14203",
        "endpoint_sensor": "1.2.0.4926"
      },
      "last_scan_time": 1733997910,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208222100",
        "536870944",
        "1208221953",
        "1208222211",
        "1208221798",
        "1208222245",
        "1208222096",
        "1208222021",
        "1090519040",
        "1208222099",
        "1073741840",
        "1073741856",
        "1208221808",
        "512",
        "1208221779",
        "1208221988",
        "536870976",
        "1208221844",
        "1208222229",
        "1208090624",
        "1208221826"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:tool_error:201"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733992246,
      "xes_ei_update_time": 1733998587,
      "epp_tracking_time": 1733988862,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "outdated",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-03",
              "name": "Extended non-use"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-04",
              "name": "Very few"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRDCOPWRDS01.armor.fr",
      "display_name": "SVFRDCOPWRDS01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "6d8dcd9f-a5d9-741b-a677-3752da9a1f13",
        "ws_group_id": "259",
        "name": "RDS",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "10",
        "name": "ARM_POL_SRV_PRO_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:1C:84",
          "ips": [
            "172.27.15.53",
            "fe80::3f9d:b293:430b:d2ee%9"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.27.15.53",
      "xdr_device_id": "79892745-a2b6-5882-a27a-73c9cf437b80",
      "ws_device_id": "2542",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734069113,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_precheck:windows:xbc_error:10001"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_precheck",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989826,
      "xes_ei_update_time": 1724902915,
      "epp_tracking_time": 1686778268,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "",
      "service_gateway_proxy_info": {
        "value": "",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 57,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-01",
              "name": "Design and development"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRDCOPDWTSQL.armor.fr",
      "display_name": "SVFRDCOPDWTSQL",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "93f0ec1e-7fa1-dd1d-f53c-678d20a441ef",
        "ws_group_id": "243",
        "name": "MSSQL",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Network adapter 1",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:D5:66",
          "ips": []
        },
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:D5:66",
          "ips": [
            "172.27.13.12"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.27.13.12",
      "xdr_device_id": "ee2b4247-e053-1db6-1f04-0a8b318e78b6",
      "ws_device_id": "4970",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734067306,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xbc_error_10001"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989941,
      "xes_ei_update_time": 1724902533,
      "epp_tracking_time": 1706528114,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 56,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "LTFRCBO676",
      "display_name": "LTFRCBO676",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "b93f8d3a-026f-4974-90bc-6b978d6af2be",
        "ws_group_id": "",
        "name": "Workgroup",
        "type": "Native",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "c1f3b6ae-1cd4-461a-9e55-d0aa910fa9f0",
        "name": "ARM_POL_COMPUTERS_ALL_MAIN"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "f4:d1:08:c4:32:31",
          "ips": [
            "172.22.25.175"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.22.25.175",
      "xdr_device_id": "1c8c1810-9608-4069-9fb5-7d4aa8a5baf3",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19045)",
        "kernel_version": ""
      },
      "cpu_arch": "64-bit",
      "last_logged_on_user": "ARMOR\\eponce",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734014698,
      "epp_component_module_status": "off",
      "epp_component_module_last_connected_time": 1734014743,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "14.0.14203",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 1733504790,
      "isolation_status": "off",
      "product_installed": [
        "sao",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221953",
        "1208222100",
        "1208221798",
        "1208222245",
        "1208222021",
        "1090519040",
        "1073741840",
        "1073741856",
        "512",
        "1208221779",
        "1208222229",
        "1208090624",
        "536870944",
        "1208222211",
        "1208222096",
        "1208222296",
        "1208222099",
        "1208221808",
        "1208221988",
        "536870976",
        "1208221844",
        "1208221826"
      ],
      "recommended_actions": {
        "errors": [
          "ncie:recording:windows:module_error:0x00000003"
        ],
        "tips": []
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "disabled",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1727250696,
      "xes_ei_update_time": 1734014759,
      "epp_tracking_time": 1723684620,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "outdated",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 49,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-01",
              "name": "Weekdays"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-01",
              "name": "Design and development"
            },
            {
              "id": "AU-03",
              "name": "Cloud services"
            },
            {
              "id": "AU-09",
              "name": "Communication"
            },
            {
              "id": "AU-11",
              "name": "Sales and marketing"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-04",
              "name": "Various"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "PCNXLWAD01.armor.fr",
      "display_name": "PCNXLWAD01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "df233f1e-b75c-6ad0-8fcb-10150b8d9058",
        "ws_group_id": "1531",
        "name": "02-NoBackup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "101",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY_ACTIVITY-MON-DISABLED"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:93:46:94",
          "ips": [
            "10.86.1.1"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.86.1.1",
      "xdr_device_id": "d8562d79-cb72-4798-b29a-a643259a8c2f",
      "ws_device_id": "4254",
      "os": {
        "name": "Microsoft Windows Server 2019",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 17763)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.2260"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "disabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125",
          "xes:fail_install:windows:xbc_error:10001"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "disabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989935,
      "xes_ei_update_time": 1733457249,
      "epp_tracking_time": 1698324397,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 62,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "IC",
          "name": "IP segment connections",
          "tags": [
            {
              "id": "IC-02",
              "name": "Few cross-segment"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        },
        {
          "id": "DR",
          "name": "Device protocols",
          "tags": [
            {
              "id": "DR-12",
              "name": "LDAP"
            },
            {
              "id": "DR-14",
              "name": "KERBEROS"
            }
          ]
        },
        {
          "id": "DS",
          "name": "Device services",
          "tags": [
            {
              "id": "DS-03",
              "name": "Kerberos authentication"
            }
          ]
        },
        {
          "id": "DC",
          "name": "Internal device connections",
          "tags": [
            {
              "id": "DC-02",
              "name": "Few"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "DTUSAMO057",
      "display_name": "",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "",
        "name": "",
        "type": "",
        "is_v1es": false,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "",
        "name": "",
        "type": ""
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "",
          "display_name": "",
          "dhcp": "unknown",
          "mac": "c4:5a:b1:fa:3e:79",
          "ips": [
            "172.25.18.149"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "[\"172.25.18.149\"]",
      "xdr_device_id": "55892f96-1b43-40b3-97e0-9f9944fe7927",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19045)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734046504,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "xbc_error_10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975324,
      "xes_ei_update_time": 1721835425,
      "epp_tracking_time": 0,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "unknown",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-02",
              "name": "Desktop"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRDCOPWAPSBK.armor.fr",
      "display_name": "SVFRDCOPWAPSBK",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "46b50daf-a531-aad8-3b84-4012fa8fd2e2",
        "ws_group_id": "289",
        "name": "APS",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:70:EE",
          "ips": [
            "172.27.12.161"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.27.12.161",
      "xdr_device_id": "feefa405-aa51-76d7-5aac-66b5f506b731",
      "ws_device_id": "5549",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734053005,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733990002,
      "xes_ei_update_time": 1731308714,
      "epp_tracking_time": 1730731925,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-01",
              "name": "High"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRDCOPWAPSCB.armor.fr",
      "display_name": "SVFRDCOPWAPSCB",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "46b50daf-a531-aad8-3b84-4012fa8fd2e2",
        "ws_group_id": "289",
        "name": "APS",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:A6:07",
          "ips": [
            "fe80::74f:80cf:d4da:b964%11",
            "172.27.12.162"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.27.12.162",
      "xdr_device_id": "0aca783c-9491-ecdf-7841-f3589c2c641b",
      "ws_device_id": "5550",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734066764,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733990360,
      "xes_ei_update_time": 1731483945,
      "epp_tracking_time": 1730732160,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRDCODWAPS.armor.fr",
      "display_name": "SVFRDCODWAPS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "46b50daf-a531-aad8-3b84-4012fa8fd2e2",
        "ws_group_id": "289",
        "name": "APS",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:BB:9E",
          "ips": [
            "172.27.12.156"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.27.12.156",
      "xdr_device_id": "d5d97d9c-6cfd-9bba-778a-111b7cf59b27",
      "ws_device_id": "5548",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734052903,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975316,
      "xes_ei_update_time": 1731483620,
      "epp_tracking_time": 1731052697,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AU",
          "name": "Primary app category",
          "tags": [
            {
              "id": "AU-10",
              "name": "Project management"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVUSHEOPWAPS.armor.fr",
      "display_name": "SVUSHEOPWAPS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "87f5c936-682e-58b1-7eb8-bc5928f7310a",
        "ws_group_id": "1630",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:9E:91",
          "ips": [
            "fe80::e379:e833:4155:e8a4%15",
            "10.1.1.13"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.1.1.13",
      "xdr_device_id": "29d0a2b1-26d2-6036-0d81-2b99156594a7",
      "ws_device_id": "5555",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734067452,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989890,
      "xes_ei_update_time": 1731516174,
      "epp_tracking_time": 1731346100,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVMXSQOPWAPS.armor.fr",
      "display_name": "SVMXSQOPWAPS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "9e6b66c0-669b-241f-017d-4922b50a7c3f",
        "ws_group_id": "1614",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:AE:FF",
          "ips": [
            "fe80::7524:97db:32bd:83f%14",
            "10.52.1.11"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.52.1.11",
      "xdr_device_id": "cd965732-42fe-09b8-b2d6-aca7d2b53708",
      "ws_device_id": "5551",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734059257,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975317,
      "xes_ei_update_time": 1731520413,
      "epp_tracking_time": 1730815251,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-03",
              "name": "Low"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVSGSGOPWAPS.armor.fr",
      "display_name": "SVSGSGOPWAPS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "b640ddf5-30c8-8d0b-9d04-7dbe9495e222",
        "ws_group_id": "1642",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:4F:91",
          "ips": [
            "10.65.1.11"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.65.1.11",
      "xdr_device_id": "452f56c0-9a87-8c8f-7e2d-015baa1e7ebd",
      "ws_device_id": "5561",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734050250,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733989854,
      "xes_ei_update_time": 1731523948,
      "epp_tracking_time": 1731429831,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVPLPROPWAPS.armor.fr",
      "display_name": "SVPLPROPWAPS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "237a17c0-de4b-e213-db5e-028fcd1428bb",
        "ws_group_id": "1609",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:1A:EA",
          "ips": [
            "fe80::1dff:2cd2:831f:fd62%13",
            "10.48.1.11"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.48.1.11",
      "xdr_device_id": "3b29be8c-a0df-28bc-924a-b5f6ae0c64ba",
      "ws_device_id": "5560",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734060821,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975317,
      "xes_ei_update_time": 1731568200,
      "epp_tracking_time": 1731425517,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVCNXLOPWKINDEE.armor.fr",
      "display_name": "SVCNXLOPWKINDEE",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "478eaf33-af67-d9bd-aa23-0c7307439a22",
        "ws_group_id": "1533",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:9D:8C",
          "ips": [
            "10.86.1.15"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.86.1.15",
      "xdr_device_id": "a96a5627-2069-a78c-9c12-3783d0a41af7",
      "ws_device_id": "5562",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222296",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1732638201,
      "xes_ei_update_time": 1733994664,
      "epp_tracking_time": 1731506311,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 56,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "svfrchopwaps.armor.fr",
      "display_name": "",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "fe48a939-1875-cecf-0fb7-37cd7f2308ea",
        "ws_group_id": "0",
        "name": "Computers",
        "type": "root",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:A9:B4:0A",
          "ips": [
            "172.30.100.12"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.30.100.12",
      "xdr_device_id": "92c904be-2444-1a5e-b3c6-c961066e0e0d",
      "ws_device_id": "5564",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734044889,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125",
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733990285,
      "xes_ei_update_time": 1731924921,
      "epp_tracking_time": 1731668613,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 58,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVUSAMOPWTEST01.armor.fr",
      "display_name": "SVUSAMOPWTEST01",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "4c01d8ea-e568-d53e-afb6-fbab91ad233a",
        "ws_group_id": "2146",
        "name": "DELL",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:77:73",
          "ips": [
            "172.25.1.100"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.25.1.100",
      "xdr_device_id": "b57e4fd4-674c-ed30-fdb8-901db32f86dd",
      "ws_device_id": "5566",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975320,
      "xes_ei_update_time": 1733808293,
      "epp_tracking_time": 1732117082,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 54,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-01",
              "name": "One regular"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "LTFRCBO798",
      "display_name": "",
      "description": "",
      "type": "desktop",
      "instance": {
        "id": "000d3a45-4275-5cb4-8c76-01fa2fdf3a54",
        "name": "Standard Endpoint Protection Manager",
        "type": "apex-central-saas",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "44740d94-245a-484c-a66f-796fa356b950",
        "name": "ARMOR",
        "type": "apex-one-saas"
      },
      "group": {
        "id": "",
        "ws_group_id": "",
        "name": "",
        "type": "",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "",
        "name": ""
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [],
      "vdi": {
        "type": ""
      },
      "last_ip_used": "[\"172.21.25.72\"]",
      "xdr_device_id": "f5e2cfe1-7814-4a9d-82fa-395211a8eb1f",
      "ws_device_id": "",
      "os": {
        "name": "Windows 10",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 19045)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "unknown",
      "last_connected_time": 0,
      "epp_component_module_status": "unknown",
      "epp_component_module_last_connected_time": 0,
      "managed": false,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "unknown",
      "product_installed": [],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "update-required"
        }
      ],
      "installed_components": [],
      "recommended_actions": {
        "errors": [
          "xbc_error_10070"
        ],
        "tips": [
          "tip-unmanaged-device-agent-install-recommended"
        ]
      },
      "xes_deployment_status": "deleted",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "unknown",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 0,
      "xes_ei_update_time": 0,
      "epp_tracking_time": 0,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "unknown",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 0,
      "risk_level": "low",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVUSAMOPWTEST02.armor.fr",
      "display_name": "SVUSAMOPWTEST02",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "4c01d8ea-e568-d53e-afb6-fbab91ad233a",
        "ws_group_id": "2146",
        "name": "DELL",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:3B:F0",
          "ips": [
            "172.25.1.101"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.25.1.101",
      "xdr_device_id": "7806613f-21e0-e514-debc-077687939fcd",
      "ws_device_id": "5567",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975321,
      "xes_ei_update_time": 1733971345,
      "epp_tracking_time": 1733472131,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 56,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-01",
              "name": "Many"
            }
          ]
        },
        {
          "id": "AL",
          "name": "Activity location",
          "tags": [
            {
              "id": "AL-03",
              "name": "Several regular"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVFRDCOPWWSUS.armor.fr",
      "display_name": "SVFRDCOPWWSUS",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "0be97f3c-8f5f-fdde-82b3-fc14f7eb69a9",
        "ws_group_id": "539",
        "name": "UPDATES",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:87:93:39",
          "ips": [
            "172.27.12.13"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "172.27.12.13",
      "xdr_device_id": "bade3263-d7b8-dd10-2256-63f7177d303b",
      "ws_device_id": "5573",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "connected",
      "last_connected_time": 1734072534,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": "1.2.0.5678"
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "enabled",
          "setting": "enabled",
          "summarized_status": "enabled"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125"
        ],
        "tips": [
          "tip-sensor-coexist-should-disable-cloud-one-workload-security-sensor"
        ]
      },
      "xes_deployment_status": "recording",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733986309,
      "xes_ei_update_time": 1733831726,
      "epp_tracking_time": 1733126685,
      "ring": "0",
      "components_update_setting": "on-schedule",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 56,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-01",
              "name": "One"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-02",
              "name": "Several"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-02",
              "name": "Shared"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVMXSQOPWSNTRN.armor.fr",
      "display_name": "SVMXSQOPWSNTRN",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "9e6b66c0-669b-241f-017d-4922b50a7c3f",
        "ws_group_id": "1614",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:44:A0",
          "ips": [
            "fe80::46f4:e4d:ce5d:d8b6%15",
            "10.52.1.18"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.52.1.18",
      "xdr_device_id": "df6a7dc1-e263-5f4a-f52c-8c6c9bde58c7",
      "ws_device_id": "5576",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734047967,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125",
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733975321,
      "xes_ei_update_time": 1733727970,
      "epp_tracking_time": 1733318584,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 54,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-03",
              "name": "Few"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    },
    {
      "hostname": "SVCATOOPWFILES.armor.fr",
      "display_name": "SVCATOOPWFILES",
      "description": "",
      "type": "server",
      "instance": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security",
        "is_v1es": true,
        "is_opt_in_version_control_policies_feature": false
      },
      "server": {
        "id": "a37c7671-f8c4-89bb-1059-5d466d574710",
        "name": "Server and Workload Protection Manager - 744503315589",
        "type": "workload-security"
      },
      "group": {
        "id": "1ff9ac89-c7d8-b011-766a-5d5956ac51e1",
        "ws_group_id": "1538",
        "name": "01-Backup",
        "type": "vcenter-folder-network",
        "xes_group_id": "",
        "xes_group_name": ""
      },
      "policy": {
        "id": "9",
        "name": "ARM_POL_SRV_ESSENTIALS_SECURITY"
      },
      "sensor_policy": "",
      "sensor_policy_override": "unknown",
      "network_interfaces": [
        {
          "name": "Ethernet0",
          "display_name": "",
          "dhcp": "false",
          "mac": "00:50:56:AA:3D:A2",
          "ips": [
            "fe80::7778:a9c5:a31d:c55%12",
            "10.11.1.6"
          ]
        }
      ],
      "vdi": {
        "type": "non-vdi"
      },
      "last_ip_used": "10.11.1.6",
      "xdr_device_id": "9d621d6b-493c-9e79-9577-76ad77ddbd9d",
      "ws_device_id": "5571",
      "os": {
        "name": "Microsoft Windows Server 2022",
        "platform": "windows",
        "arch": "x86_64",
        "version": "10.0 (Build 20348)",
        "kernel_version": ""
      },
      "cpu_arch": "unknown",
      "last_logged_on_user": "",
      "endpoint_connectivity": "disconnected",
      "last_connected_time": 1734053645,
      "epp_component_module_status": "on",
      "epp_component_module_last_connected_time": 1734072587,
      "managed": true,
      "tags": [],
      "agent_version": {
        "endpoint_protection": "20.0.1.23340",
        "endpoint_sensor": ""
      },
      "last_scan_time": 0,
      "isolation_status": "off",
      "product_installed": [
        "sds",
        "xes"
      ],
      "xes_features": [
        {
          "name": "endpoint-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "enabled"
        },
        {
          "name": "advanced-risk-telemetry",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        },
        {
          "name": "data-sensor",
          "status": "unknown",
          "setting": "unknown",
          "summarized_status": "unknown"
        }
      ],
      "installed_components": [
        "1082130432",
        "1208221992",
        "1208221733",
        "1208222001",
        "1208221798",
        "1208221732",
        "1208222289",
        "1090519040",
        "1073741840",
        "2048",
        "1208221779",
        "1208221976",
        "1208090624",
        "1208221762",
        "1208222214",
        "1208222019",
        "536870944",
        "1208221764",
        "536871936",
        "1208221763",
        "1208222099",
        "1208221801",
        "536870976",
        "1208221800"
      ],
      "recommended_actions": {
        "errors": [
          "inspect_error_c1_detected_125",
          "xes:fail_install:windows:xbc_error:10070"
        ],
        "tips": []
      },
      "xes_deployment_status": "fail_install",
      "integrated_endpoint_sensor_status": "unknown",
      "activity_monitoring_status": "enabled",
      "unmanaged_discovered_instance": "",
      "unmanaged_discovered_source": "",
      "pending_changes": [],
      "xes_policy_update_time": 1733986806,
      "xes_ei_update_time": 1733829266,
      "epp_tracking_time": 1732725369,
      "ring": "",
      "components_update_setting": "not-supported",
      "components_update_status": "latest",
      "service_gateway_proxy": "Direct connect",
      "service_gateway_proxy_info": {
        "value": "Direct connect",
        "fqdn": "",
        "trimmed": false,
        "ips": []
      },
      "risk_score": 56,
      "risk_level": "medium",
      "security_update_plan": "",
      "agent_version_status": "",
      "agent_update_policy": "",
      "program_update_status": "notSupported",
      "custom_tags_count": 0,
      "data_sensor_status": "unknown",
      "data_sensor_setting": "unknown",
      "predefined_tags": [
        {
          "id": "XH",
          "name": "Account/Device access",
          "tags": [
            {
              "id": "XH-02",
              "name": "Regular access"
            }
          ]
        },
        {
          "id": "DE",
          "name": "Average activity detections",
          "tags": [
            {
              "id": "DE-02",
              "name": "Medium"
            }
          ]
        },
        {
          "id": "DVT",
          "name": "Device type",
          "tags": [
            {
              "id": "DVT-01",
              "name": "Server"
            }
          ]
        },
        {
          "id": "AD",
          "name": "Activity occurrence",
          "tags": [
            {
              "id": "AD-03",
              "name": "Whole week"
            }
          ]
        },
        {
          "id": "NA",
          "name": "IP addresses",
          "tags": [
            {
              "id": "NA-02",
              "name": "Few"
            }
          ]
        },
        {
          "id": "MA",
          "name": "Active days",
          "tags": [
            {
              "id": "MA-02",
              "name": "Several"
            }
          ]
        },
        {
          "id": "DO",
          "name": "Device usage",
          "tags": [
            {
              "id": "DO-01",
              "name": "Exclusive"
            }
          ]
        },
        {
          "id": "DP",
          "name": "Device operating system",
          "tags": [
            {
              "id": "DP-03",
              "name": "Windows"
            }
          ]
        }
      ],
      "is_v1_endpoint_deleting": false
    }
  ],
  "total": 50
}`) // 通常在code中，會用「原始字串 (``)」來呈現JSON string

	var apiResult map[string]interface{}

	if err := json.Unmarshal(json_input, &apiResult); err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		for k, v := range apiResult {
			if k == "endpoints" {
				if endpoints, ok := v.([]interface{}); ok {
					for _, endpoint := range endpoints {
						if endpoint, ok := endpoint.(map[string]interface{}); ok {
							for k2, v2 := range endpoint {
								if k2 == "hostname" {
									fmt.Printf("%v\n", v2)
								}
							}
						}
					}
				}
			}
		}
	}
}
