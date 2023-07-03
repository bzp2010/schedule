import { ApolloServer } from "@apollo/server";
import { GraphQLFileLoader } from "@graphql-tools/graphql-file-loader";
import { loadTypedefs } from "@graphql-tools/load";

import { resolvers } from "@/app/api/graphql/resolvers";
import { EntityManager, MikroORM } from "@mikro-orm/core";

export interface ApolloServerContext {
  db: MikroORM
  em: EntityManager
}

const typeDefs = (await loadTypedefs("**/*.graphql", {
  loaders: [new GraphQLFileLoader()],
  noRequire: true,
})).map(src => src.document ?? "");

const server = new ApolloServer<ApolloServerContext>({
  typeDefs: typeDefs,
  resolvers: resolvers,
});

export default server;
