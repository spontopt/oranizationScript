package process

import (
	"encoding/json"
	"log"
	"organizationScript/data"
)

func ImportStatus(db data.MongoDb) {
	uo := db.GetAllUserOrganizations()

	subscriptions := make(map[string][]data.UserOrganization)
	for _, o := range uo {
		subscriptions[o.SubscriptionID] = append(subscriptions[o.SubscriptionID], o)
	}

	for s, uos := range subscriptions {
		users := db.GetUsers(s)
		for _, u := range users {
			userOrganizations := make([]data.UserOrganization, 0)
			for _, uo := range uos {
				if uo.UserID == u.IDUser {
					uo.Status = u.Active
					role := uo.RoleID
					if role == "SuperAdmin" {
						role = "Super Admin"
					}
					uo.Role = role

					userOrganizations = append(userOrganizations, uo)
				}
			}
			a, _ := json.Marshal(userOrganizations)
			log.Println("users:", string(a))
			db.BulkSaveUserOrganizations(userOrganizations)
		}
	}

}
