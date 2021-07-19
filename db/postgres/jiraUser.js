'use strict'

const {
  Model
} = require('sequelize')

module.exports = (sequelize, DataTypes) => {
  class JiraUser extends Model {

  }

  JiraUser.init({
    uuid: {
      primaryKey: true,
      type: DataTypes.UUID,
      defaultValue: DataTypes.UUIDV4
    },
    self: DataTypes.STRING,
    accountId: DataTypes.STRING,
    name: DataTypes.STRING,
    key: DataTypes.STRING,
    emailAddress: DataTypes.STRING,
    displayName: DataTypes.STRING,
    active: DataTypes.BOOLEAN,
    timezone: DataTypes.STRING,
    createdAt: {
      allowNull: false,
      type: DataTypes.DATE,
      defaultValue: DataTypes.NOW
    },
    updatedAt: {
      allowNull: false,
      type: DataTypes.DATE,
      defaultValue: DataTypes.NOW
    }
  }, {
    sequelize,
    modelName: 'JiraUser',
    underscored: true
  })

  return JiraUser
}
