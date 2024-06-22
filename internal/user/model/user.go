package model

import "time"

type User struct {
	CompanyId   string     `yaml:"company_id" mapstructure:"company_id" json:"companyId" gorm:"column:company_id;primary_key" dynamodbav:"companyId" firestore:"-" avro:"companyId" validate:"required,max=40" operator:"="`
	UserId      string     `yaml:"user_id" mapstructure:"user_id" json:"userId" gorm:"column:user_id;primary_key" dynamodbav:"userId" firestore:"-" avro:"userId" validate:"required,max=40" operator:"="`
	Username    string     `yaml:"username" mapstructure:"username" json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" avro:"username" validate:"required,username,max=100"`
	Email       string     `yaml:"email" mapstructure:"email" json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" avro:"email" validate:"email,max=100"`
	Phone       string     `yaml:"phone" mapstructure:"phone" json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" avro:"phone" validate:"required,phone,max=18" operator:"like"`
	DateOfBirth *time.Time `yaml:"date_of_birth" mapstructure:"date_of_birth" json:"dateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth" avro:"dateOfBirth"`
}
