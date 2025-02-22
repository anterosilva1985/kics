package Cx

import data.generic.common as common_lib

CxPolicy[result] {
	doc := input.document[i]

	[path, value] = walk(doc)

	value.type == "Microsoft.Compute/virtualMachines"
	not is_windows(value)
	not value.properties.osProfile.linuxConfiguration.disablePasswordAuthentication

	issue := prepare_issue(value)

	result := {
		"documentId": input.document[i].id,
		"resourceType": value.type,
		"resourceName": value.name,
		"searchKey": sprintf("%s.name=%s%s", [common_lib.concat_path(path), value.name, issue.sk]),
		"issueType": issue.issueType,
		"keyExpectedValue": "'disablePasswordAuthentication' should be set to true",
		"keyActualValue": issue.keyActualValue,
		"searchLine": common_lib.build_search_line(path, issue.sl),
	}
}

is_windows(resource) {
	contains(lower(resource.properties.storageProfile.imageReference.publisher), "windows")
}

prepare_issue(resource) = issue {
	resource.properties.osProfile.linuxConfiguration.disablePasswordAuthentication == false
	issue := {
		"resourceType": resource.type,
		"resourceName": resource.name,
		"issueType": "IncorrectValue",
		"keyActualValue": "'disablePasswordAuthentication' is set to false",
		"sk": ".properties.osProfile.linuxConfiguration.disablePasswordAuthentication",
		"sl": ["properties", "osProfile", "linuxConfiguration", "disablePasswordAuthentication"],
	}
} else = issue {
	issue := {
		"resourceType": resource.type,
		"resourceName": resource.name,
		"issueType": "MissingAttribute",
		"keyActualValue": "'disablePasswordAuthentication' is undefined",
		"sk": "",
		"sl": ["name"],

	}
}
