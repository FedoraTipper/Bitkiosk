import { GraphQLClient } from "graphql-request";
const config = require("@/utils/config/config");

export default class GQLClientFactory {
  GQL_ENDPOINT = config.default.BASE_PATH + config.default.PATH.GRAPHQL;
  constructor() {}

  newGQLClient(): GraphQLClient {
    const newGQLClient = new GraphQLClient(this.GQL_ENDPOINT, {
      credentials: "include"
    });

    return newGQLClient;
  }
}
