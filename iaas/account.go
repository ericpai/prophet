package iaas

func ListAccounts() map[string][]string {
	result := make(map[string][]string)
	for name, account := range prophetConfig.Accounts {
		if _, exist := result[account.Provider]; !exist {
			result[account.Provider] = make([]string, 0)
		}
		result[account.Provider] = append(result[account.Provider], name)
	}
	return result
}
