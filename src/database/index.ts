import { EntityManager, MikroORM } from "@mikro-orm/core";
import { SqliteDriver } from "@mikro-orm/sqlite";
import { TsMorphMetadataProvider } from "@mikro-orm/reflection";
import { Task } from "./entities/task";
import config from "@/mikro-orm.config";

const db = {
  getORM: async () => {
    return await MikroORM.init<SqliteDriver>(config);
  },
};

class Datebase {
  // private 只允许在类内访问
  private static instance: MikroORM | null = null;
  // 获取单例
  static async getORM(): Promise<MikroORM> {
    if (Datebase.instance === null) {
      Datebase.instance = await MikroORM.init<SqliteDriver>(config);
    }
    return Datebase.instance;
  }
  static getEM(): EntityManager {
    return (Datebase.instance as unknown as MikroORM).em.fork();
  }
}

export default Datebase;
