// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingIp6ConfigKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_IP6_CONFIG_METHOD:
		t = ktypeString
	case NM_SETTING_IP6_CONFIG_DNS:
		t = ktypeWrapperIpv6Dns
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		t = ktypeArrayString
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		t = ktypeWrapperIpv6Addresses
	case NM_SETTING_IP6_CONFIG_ROUTES:
		t = ktypeWrapperIpv6Routes
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		t = ktypeBoolean
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		t = ktypeInt32
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		t = ktypeString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingIp6Config(key string) bool {
	switch key {
	case NM_SETTING_IP6_CONFIG_METHOD:
		return true
	case NM_SETTING_IP6_CONFIG_DNS:
		return true
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		return true
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		return true
	case NM_SETTING_IP6_CONFIG_ROUTES:
		return true
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		return true
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		return true
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		return true
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		return true
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		return true
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		return true
	}
	return false
}

// Get key's default value
func getSettingIp6ConfigKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		Logger.Error("invalid key:", key)
	case NM_SETTING_IP6_CONFIG_METHOD:
		valueJSON = `""`
	case NM_SETTING_IP6_CONFIG_DNS:
		valueJSON = `null`
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		valueJSON = `null`
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		valueJSON = `null`
	case NM_SETTING_IP6_CONFIG_ROUTES:
		valueJSON = `null`
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		valueJSON = `false`
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		valueJSON = `false`
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		valueJSON = `false`
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		valueJSON = `true`
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		valueJSON = `-1`
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		valueJSON = `""`
	}
	return
}

// Get JSON value generally
func generalGetSettingIp6ConfigKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		Logger.Error("generalGetSettingIp6ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP6_CONFIG_METHOD:
		value = getSettingIp6ConfigMethodJSON(data)
	case NM_SETTING_IP6_CONFIG_DNS:
		value = getSettingIp6ConfigDnsJSON(data)
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		value = getSettingIp6ConfigDnsSearchJSON(data)
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		value = getSettingIp6ConfigAddressesJSON(data)
	case NM_SETTING_IP6_CONFIG_ROUTES:
		value = getSettingIp6ConfigRoutesJSON(data)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		value = getSettingIp6ConfigIgnoreAutoRoutesJSON(data)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		value = getSettingIp6ConfigIgnoreAutoDnsJSON(data)
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		value = getSettingIp6ConfigNeverDefaultJSON(data)
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		value = getSettingIp6ConfigMayFailJSON(data)
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		value = getSettingIp6ConfigIp6PrivacyJSON(data)
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		value = getSettingIp6ConfigDhcpHostnameJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingIp6ConfigKeyJSON(data _ConnectionData, key, valueJSON string) {
	switch key {
	default:
		Logger.Error("generalSetSettingIp6ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP6_CONFIG_METHOD:
		logicSetSettingIp6ConfigMethodJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_DNS:
		setSettingIp6ConfigDnsJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_DNS_SEARCH:
		setSettingIp6ConfigDnsSearchJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_ADDRESSES:
		setSettingIp6ConfigAddressesJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_ROUTES:
		setSettingIp6ConfigRoutesJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES:
		setSettingIp6ConfigIgnoreAutoRoutesJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS:
		setSettingIp6ConfigIgnoreAutoDnsJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_NEVER_DEFAULT:
		setSettingIp6ConfigNeverDefaultJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_MAY_FAIL:
		setSettingIp6ConfigMayFailJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_IP6_PRIVACY:
		setSettingIp6ConfigIp6PrivacyJSON(data, valueJSON)
	case NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME:
		setSettingIp6ConfigDhcpHostnameJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingIp6ConfigMethodExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD)
}
func isSettingIp6ConfigDnsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS)
}
func isSettingIp6ConfigDnsSearchExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH)
}
func isSettingIp6ConfigAddressesExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES)
}
func isSettingIp6ConfigRoutesExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES)
}
func isSettingIp6ConfigIgnoreAutoRoutesExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES)
}
func isSettingIp6ConfigIgnoreAutoDnsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS)
}
func isSettingIp6ConfigNeverDefaultExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT)
}
func isSettingIp6ConfigMayFailExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL)
}
func isSettingIp6ConfigIp6PrivacyExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY)
}
func isSettingIp6ConfigDhcpHostnameExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME)
}

// Ensure field and key exists and not empty
func ensureFieldSettingIp6ConfigExists(data _ConnectionData, errs map[string]string, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_IP6_CONFIG_SETTING_NAME) {
		rememberError(errs, relatedKey, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_IP6_CONFIG_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_IP6_CONFIG_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_IP6_CONFIG_SETTING_NAME))
	}
}
func ensureSettingIp6ConfigMethodNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigMethodExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_METHOD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigMethod(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_METHOD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigDnsNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigDnsExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigDns(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DNS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigDnsSearchNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigDnsSearchExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DNS_SEARCH, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigDnsSearch(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DNS_SEARCH, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigAddressesNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigAddressesExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_ADDRESSES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigAddresses(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_ADDRESSES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigRoutesNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigRoutes(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_ROUTES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp6ConfigIgnoreAutoRoutesNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigIgnoreAutoRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigIgnoreAutoDnsNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigIgnoreAutoDnsExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigNeverDefaultNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigNeverDefaultExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigMayFailNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigMayFailExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_MAY_FAIL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigIp6PrivacyNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigIp6PrivacyExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp6ConfigDhcpHostnameNoEmpty(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigDhcpHostnameExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp6ConfigDhcpHostname(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingIp6ConfigMethod(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD).(string)
	return
}
func getSettingIp6ConfigDns(data _ConnectionData) (value [][]byte) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS).([][]byte)
	return
}
func getSettingIp6ConfigDnsSearch(data _ConnectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH).([]string)
	return
}
func getSettingIp6ConfigAddresses(data _ConnectionData) (value Ipv6Addresses) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES).(Ipv6Addresses)
	return
}
func getSettingIp6ConfigRoutes(data _ConnectionData) (value Ipv6Routes) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES).(Ipv6Routes)
	return
}
func getSettingIp6ConfigIgnoreAutoRoutes(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES).(bool)
	return
}
func getSettingIp6ConfigIgnoreAutoDns(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS).(bool)
	return
}
func getSettingIp6ConfigNeverDefault(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT).(bool)
	return
}
func getSettingIp6ConfigMayFail(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL).(bool)
	return
}
func getSettingIp6ConfigIp6Privacy(data _ConnectionData) (value int32) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY).(int32)
	return
}
func getSettingIp6ConfigDhcpHostname(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME).(string)
	return
}

// Setter
func setSettingIp6ConfigMethod(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, value)
}
func setSettingIp6ConfigDns(data _ConnectionData, value [][]byte) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, value)
}
func setSettingIp6ConfigDnsSearch(data _ConnectionData, value []string) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, value)
}
func setSettingIp6ConfigAddresses(data _ConnectionData, value Ipv6Addresses) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, value)
}
func setSettingIp6ConfigRoutes(data _ConnectionData, value Ipv6Routes) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, value)
}
func setSettingIp6ConfigIgnoreAutoRoutes(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, value)
}
func setSettingIp6ConfigIgnoreAutoDns(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, value)
}
func setSettingIp6ConfigNeverDefault(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, value)
}
func setSettingIp6ConfigMayFail(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, value)
}
func setSettingIp6ConfigIp6Privacy(data _ConnectionData, value int32) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, value)
}
func setSettingIp6ConfigDhcpHostname(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, value)
}

// JSON Getter
func getSettingIp6ConfigMethodJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_METHOD))
	return
}
func getSettingIp6ConfigDnsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS))
	return
}
func getSettingIp6ConfigDnsSearchJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS_SEARCH))
	return
}
func getSettingIp6ConfigAddressesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ADDRESSES))
	return
}
func getSettingIp6ConfigRoutesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ROUTES))
	return
}
func getSettingIp6ConfigIgnoreAutoRoutesJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES))
	return
}
func getSettingIp6ConfigIgnoreAutoDnsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS))
	return
}
func getSettingIp6ConfigNeverDefaultJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_NEVER_DEFAULT))
	return
}
func getSettingIp6ConfigMayFailJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_MAY_FAIL))
	return
}
func getSettingIp6ConfigIp6PrivacyJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IP6_PRIVACY))
	return
}
func getSettingIp6ConfigDhcpHostnameJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME))
	return
}

// JSON Setter
func setSettingIp6ConfigMethodJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_METHOD))
}
func setSettingIp6ConfigDnsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS))
}
func setSettingIp6ConfigDnsSearchJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DNS_SEARCH))
}
func setSettingIp6ConfigAddressesJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ADDRESSES))
}
func setSettingIp6ConfigRoutesJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_ROUTES))
}
func setSettingIp6ConfigIgnoreAutoRoutesJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES))
}
func setSettingIp6ConfigIgnoreAutoDnsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS))
}
func setSettingIp6ConfigNeverDefaultJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_NEVER_DEFAULT))
}
func setSettingIp6ConfigMayFailJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_MAY_FAIL))
}
func setSettingIp6ConfigIp6PrivacyJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_IP6_PRIVACY))
}
func setSettingIp6ConfigDhcpHostnameJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME, valueJSON, getSettingIp6ConfigKeyType(NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME))
}

// Remover
func removeSettingIp6ConfigMethod(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD)
}
func removeSettingIp6ConfigDns(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS)
}
func removeSettingIp6ConfigDnsSearch(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS_SEARCH)
}
func removeSettingIp6ConfigAddresses(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES)
}
func removeSettingIp6ConfigRoutes(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES)
}
func removeSettingIp6ConfigIgnoreAutoRoutes(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES)
}
func removeSettingIp6ConfigIgnoreAutoDns(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS)
}
func removeSettingIp6ConfigNeverDefault(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_NEVER_DEFAULT)
}
func removeSettingIp6ConfigMayFail(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_MAY_FAIL)
}
func removeSettingIp6ConfigIp6Privacy(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_IP6_PRIVACY)
}
func removeSettingIp6ConfigDhcpHostname(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME)
}