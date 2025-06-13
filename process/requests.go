package process

import (
	"log"
	"organizationScript/data"
	"strconv"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func requestPage(page string) data.SubscriptionsResponse {

	client := resty.New()

	subscriptions := data.SubscriptionsResponse{}
	_, err := client.R().
		SetResult(&subscriptions).
		SetHeader("x-api-key", data.KiaiHubApiKey).
		SetQueryParam("product", "launchsoft").
		SetQueryParam("limit", "100").
		SetQueryParam("offset", page).
		Get(data.KiaiHubUrl + "/api/subscriptions")

	if err != nil {
		panic(err)
	}

	return subscriptions

}

func GetAllSubscriptions(db data.MongoDb) {

	subscriptions := make([]data.Subscription, 0)
	counter := 0
	page := "0"
	end := false
	for !end {
		counter++

		log.Println("requesting page: " + page)
		s := requestPage(page)
		subscriptions = append(subscriptions, s.Subscriptions...)
		page = strconv.Itoa(s.Meta.NextPage)

		log.Println("started processing" + page)
		ProcessSubscriptions(subscriptions, db)

		log.Println("request n " + strconv.Itoa(counter))
		if s.Meta.Count < 100 {
			end = true
		}

	}
}

func ProcessSubscriptions(subscriptions []data.Subscription, db data.MongoDb) {

	for _, s := range subscriptions {
		log.Println("processing subscription: " + s.Id)
		log.Println("init roles")

		roles := initRoles()
		log.Println("get users")

		users := db.GetUsers(s.Id)
		log.Println("convert organization")

		organization := convertToOrganization(s, users)
		log.Println("saving organization")

		o := db.SaveOrganization(organization)
		log.Println("saving roles")

		db.SaveRoles(roles, o.OrganizationID.Hex(), s.Id)
		onlyUser := len(users) == 1
		userOrganizations := make([]data.UserOrganization, 0)
		log.Println("converting user organizations")

		for _, u := range users {
			userOrganizations = append(userOrganizations, convertToUserOrganization(u, o.OrganizationID.Hex(), onlyUser))
		}
		log.Println("saving user organizations")

		db.BulkSaveUserOrganizations(userOrganizations)

	}

}

func convertToOrganization(subscription data.Subscription, users []data.IUser) data.Organization {

	organization := data.Organization{
		SubscriptionID: subscription.Id,
		Email:          subscription.Email,
		Language:       "pt",
		Currency:       "EUR",
		NumberFormat:   "1.234,56",
		DateFormat:     "dd/MM/yyyy",
	}
	for _, u := range users {
		if u.Email == subscription.Email {
			return data.Organization{
				SubscriptionID: subscription.Id,
				Email:          u.Email,
				Language:       u.Language,
				Currency:       u.Currency,
				NumberFormat:   u.NumberFormat,
				DateFormat:     u.DateFormat,
			}
		}

		if u.Profile.LegacyID == "SuperAdmin" {
			organization = data.Organization{
				SubscriptionID: subscription.Id,
				Email:          subscription.Email,
				Language:       u.Language,
				Currency:       u.Currency,
				NumberFormat:   u.NumberFormat,
				DateFormat:     u.DateFormat,
			}

		}

	}
	return organization
}

func convertToUserOrganization(user data.IUser, organization string, onlyUser bool) data.UserOrganization {
	profile := "Admin"
	pName := "Admin"
	if onlyUser {
		profile = "SuperAdmin"
		pName = "Super Admin"
	}
	if user.Profile.LegacyID != "" {
		profile = user.Profile.LegacyID
		pName = user.Profile.Profile
	}

	return data.UserOrganization{
		UserOrganizationID: primitive.NilObjectID,
		OrganizationID:     organization,
		SubscriptionID:     user.IDTenant,
		RoleID:             profile,
		UserID:             user.IDUser,
		Name:               user.Name,
		Email:              user.Email,
		Role:               pName,
		Status:             user.Active,
	}
}

func initRoles() []data.Profile {
	return []data.Profile{
		{
			Profile:        "Super Admin",
			LegacyID:       "SuperAdmin",
			DefaultProfile: true,
			Custom:         false,
			Description:    "SUPERADMINDESC",
			Permissions: data.Permission{
				Features: []data.Feature{
					{
						Feature: "Dashboards",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    true,
							},
						},
					},
					{
						Feature: "Integrations",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    true,
							},
						},
					}, {
						Feature: "Users",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    true,
							},
						},
					}, {
						Feature: "Insights",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
						},
					}, {
						Feature: "Subscriptions",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
						},
					}, {
						Feature: "Permissions",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
						},
					},
				},
			},
		},
		{
			Profile:        "Admin",
			LegacyID:       "Admin",
			DefaultProfile: true,
			Custom:         false,
			Description:    "ADMINDESC",
			Permissions: data.Permission{
				Features: []data.Feature{
					{
						Feature: "Dashboards",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    true,
							},
						},
					},
					{
						Feature: "Integrations",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    true,
							},
						},
					}, {
						Feature: "Users",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    true,
							},
						},
					}, {
						Feature: "Insights",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
						},
					}, {
						Feature: "Subscriptions",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    false,
							},
						},
					}, {
						Feature: "Permissions",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    false,
							},
						},
					},
				},
			},
		},
		{
			Profile:        "User",
			LegacyID:       "User",
			DefaultProfile: true,
			Custom:         false,
			Description:    "ESERDESC",
			Permissions: data.Permission{
				Features: []data.Feature{
					{
						Feature: "Dashboards",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    false,
							},
						},
					},
					{
						Feature: "Integrations",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
							{
								Operation: "Create",
								Active:    true,
							},
							{
								Operation: "Edit",
								Active:    true,
							},
							{
								Operation: "Delete",
								Active:    false,
							},
						},
					}, {
						Feature: "Users",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    false,
							},
							{
								Operation: "Create",
								Active:    false,
							},
							{
								Operation: "Edit",
								Active:    false,
							},
							{
								Operation: "Delete",
								Active:    false,
							},
						},
					}, {
						Feature: "Insights",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    true,
							},
						},
					}, {
						Feature: "Subscriptions",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    false,
							},
						},
					}, {
						Feature: "Permissions",
						Operations: []data.Operation{
							{
								Operation: "Access",
								Active:    false,
							},
						},
					},
				},
			},
		},
	}

}
