package api_test

import (

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)


var _ = Describe("Api", func (){
	var testDB db.DB
	var userRepo repository.UserRepository
	var server api.API

	BeforeEach(func (){
		tables := map[db.TableName]db.Rows{
			"users": {
				db.Row{"username", "password"},
				db.Row{?, ?},
			},
			"kampus": {
				db.Row{"name", "email", "info", "jurusan1", "jurusan2"},
				db.Row{?, ?, ?, ?, ?},
			},
			"review": {
				db.Row{"username", "email", "nama_kampus", "jurusan", "isian"},
			}
		}
		testDB = db.NewDB(tables)

		userRepo = repository.NewUserRepository(testDB)
		kampusRepo = repository.NewKampusRepository(testDB)
		reviewRepo = repository.NewReviewRepository(testDB)

		server = api.NewAPI(userRepo, kampusRepo, reviewRepo)
	})

	Describe("/users", func (){
		It("should return a list of users", func (){
			users := server.GetUsers()
			Expect(users).To(HaveLen(2))
		})
	})
	Describe("/kampus", func (){
		It("should return a list of kampus", func (){
			kampus := server.GetKampus()
			Expect(kampus).To(HaveLen(2))
		})
	})
	Describe("/review", func (){
		It("should return a list of review", func (){
			review := server.GetReview()
			Expect(review).To(HaveLen(1))
		})
	})
	when("/users/{username}", func (){
		It("should return a user", func (){
			user := server.GetUser("username")
			Expect(user).To(Equal(db.Row{"username", "password"}))
		})
		
	})
}())