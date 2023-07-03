import { defineConfig } from "@mikro-orm/core";
import { SqliteDriver } from "@mikro-orm/sqlite";
import { Task, TaskRule } from "@/database/entities";
import { Job } from "./database/entities/job";

export default defineConfig<SqliteDriver>({
  entities: [Task, TaskRule, Job],
  dbName: "dev.db",
  type: "sqlite",
  debug: true,
  cache: {
    enabled: true,
  } 
});
