import { startServerAndCreateNextHandler } from "@as-integrations/next";
import { NextRequest } from "next/server";
import server, { ApolloServerContext } from "@/app/api/graphql/server";
import db from "@/database";

class handler {
  private static instance: ((req: Request | NextRequest, res?: undefined) => Promise<Response>) | null = null;
  static get() {
    if (handler.instance === null) {
      handler.instance = startServerAndCreateNextHandler<NextRequest, ApolloServerContext>(server, {
        context: async (req) => ({ req, db: await db.getORM(), em: db.getEM() }),
      });
    }
    return handler.instance;
  }
}

export async function GET(request: NextRequest) {
  return handler.get()(request);
}

export async function POST(request: NextRequest) {
  return handler.get()(request);
}
