package bank_example_test

import (
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bank-example/internal/models"
)

var _ = Describe("Users", func() {
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var dbm *sql.DB
		var err error

		dbm, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err := gorm.Open("postgres", dbm)
		Expect(err).ShouldNot(HaveOccurred())

		models.SetDB(gdb)
	})

	Context("fetch user", func() {
		It("found", func() {
			user := &models.Users{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				Email:     "test@test.com",
			}

			rows := sqlmock.
				NewRows([]string{"id", "first_name", "last_name", "email"}).
				AddRow(user.ID, user.FirstName, user.LastName, user.Email)

			const sqlSelectOne = `SELECT * FROM "users WHERE "users"."id" = ? LIMIT 1"`

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
				WithArgs(user.ID).
				WillReturnRows(rows)

			dbUser := user.Get(user.ID)
			dbUserID := dbUser["user"].(*models.Users).ID
			Expect(dbUserID).Should(Equal(user.ID))
		})

		It("not found", func() {
			user := &models.Users{}
			// ignore sql match
			mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
			dbUser := user.Get(2)
			Expect(len(dbUser)).Should(Equal(0))
		})
	})
})
