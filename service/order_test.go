package service

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"orderManagement/initialize"
	"orderManagement/models"
	"reflect"
	"regexp"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	mock, db := mockDB(t)

	defer db.Close()

	type args struct {
		order models.Order
	}

	tests := []struct {
		name string
		args args
		want models.OrderResponse
	}{
		{
			name: "test1",
			args: args{
				order: models.Order{
					Id:           "12345",
					Status:       "pending",
					Total:        12,
					CurrencyUnit: "INR",
					Items: []models.Item{
						{
							Id:          "54321",
							Description: "eye",
							Price:       15,
							Quantity:    2,
							OrderId:     "12345",
						},
					},
				},
			},
			want: models.OrderResponse{
				Id:           "12345",
				Status:       "pending",
				Total:        12,
				CurrencyUnit: "INR",
				Items: []models.ItemResponse{
					{
						Id:          "54321",
						Description: "eye",
						Price:       15,
						Quantity:    2,
					},
				},
			},
		},
	}

	mock.ExpectBegin()

	mock.ExpectExec("INSERT").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateOrder(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockDB(t *testing.T) (sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := mysql.New(mysql.Config{
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	initialize.DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return mock, db
}

func TestGetOrder(t *testing.T) {
	mock, db := mockDB(t)

	defer db.Close()

	type args struct {
		order models.FilterOrd
	}
	tests := []struct {
		name string
		args args
		want models.CreateOrderResponse
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				order: models.FilterOrd{
					Id:           "12345",
					Status:       "pending",
					Total:        12,
					CurrencyUnit: "INR",
				},
			},
			want: models.CreateOrderResponse{
				OrderResponse: []models.OrderResponse{
					{
						Id:           "12345",
						Status:       "pending",
						Total:        12,
						CurrencyUnit: "INR",
						Items: []models.ItemResponse{
							{
								Id:          "54321",
								Description: "eye",
								Price:       15,
								Quantity:    2,
							},
						},
					},
				},
			},
		},
	}

	selectQuery := "SELECT * FROM `orders` WHERE `orders`.`id` = ? AND `orders`.`status` = ? AND `orders`.`total` = ? AND `orders`.`currency_unit` = ? AND `orders`.`deleted_at` IS NULL"

	row := sqlmock.NewRows([]string{"id", "status", "total", "currency_unit"}).AddRow("12345", "pending", 12.0, "INR")

	mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).WithArgs("12345", "pending", 12.0, "INR").WillReturnRows(row)

	selectQuery = "SELECT * FROM `items` WHERE `items`.`order_id` = ? AND `items`.`deleted_at` IS NULL"

	row = sqlmock.NewRows([]string{"id", "description", "price", "quantity", "order_id"}).AddRow("54321", "eye", 15, 2, "12345")

	mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).WithArgs("12345").WillReturnRows(row)
	//mock.ExpectCommit()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrder(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	mock, db := mockDB(t)

	defer db.Close()

	type args struct {
		order   models.UpdateOrder
		orderId string
	}
	tests := []struct {
		name string
		args args
		want models.OrderResponse
	}{
		// TODO: Add test cases.{
		{
			name: "test1",
			args: args{
				order: models.UpdateOrder{
					Status:       "pending",
					Total:        15,
					CurrencyUnit: "INR",
				},
				orderId: "12345",
			},
			want: models.OrderResponse{
				Id:           "12345",
				Status:       "pending",
				Total:        15,
				CurrencyUnit: "INR",
				Items: []models.ItemResponse{
					{
						Id:          "54321",
						Description: "eye",
						Price:       15,
						Quantity:    2,
					},
				},
			},
		},
	}

	mock.ExpectBegin()

	selectQuery := "UPDATE `orders` SET `id`=?,`updated_at`=?,`status`=?,`total`=?,`currency_unit`=? WHERE `orders`.`deleted_at` IS NULL AND `id` = ?"

	mock.ExpectExec(regexp.QuoteMeta(selectQuery)).WithArgs("12345", sqlmock.AnyArg(), "pending", 15.0, "INR", "12345").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	selectQuery = "SELECT * FROM `orders` WHERE id=? AND `orders`.`deleted_at` IS NULL AND `orders`.`id` = ? ORDER BY `orders`.`id` LIMIT 1"

	row := sqlmock.NewRows([]string{"id", "status", "total", "currency_unit"}).AddRow("12345", "pending", 15.0, "INR")

	mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).WithArgs("12345", "12345").WillReturnRows(row)

	selectQuery = "SELECT * FROM `items` WHERE `items`.`order_id` = ? AND `items`.`deleted_at` IS NULL"

	row = sqlmock.NewRows([]string{"id", "description", "price", "quantity", "order_id"}).AddRow("54321", "eye", 15, 2, "12345")

	mock.ExpectQuery(regexp.QuoteMeta(selectQuery)).WithArgs("12345").WillReturnRows(row)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateOrder(tt.args.order, tt.args.orderId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
