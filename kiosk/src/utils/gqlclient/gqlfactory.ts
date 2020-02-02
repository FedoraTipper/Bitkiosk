import { GraphQLClient } from "graphql-request";
import CookieUtil from "@/utils/cookie/cookieutil";
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
