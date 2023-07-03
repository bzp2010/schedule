import db from "@/database";
import { TaskRule } from "@/database/entities";
import { Job } from "@/database/entities/job";
import { Task, TaskType } from "@/database/entities/task";
import { NextResponse } from "next/server";

export async function GET() {
  const orm = await db.getORM();
  //await orm.getSchemaGenerator().createDatabase("test");
  await orm.getSchemaGenerator().dropSchema();
  await orm.getSchemaGenerator().createSchema();

  const task = new Task("TaskName_1", TaskType.Shell, {
    command: "ls",
    timeout: 1
  });
  task.lastRunningAt = new Date;
  task.lastRunningTime = 1000;
  const tr = new TaskRule(task, "* * * * *", "RULE1");
  task.rules.add(tr, new TaskRule(task, "* * * * *", "RULE2"));
  const j = new Job(task, tr, new Date, new Date);

  await orm.em.fork().persistAndFlush([task, j]);

  /* const g = await db.getEM().findOne(Task, {
    configuration: {
      command: "ls"
    }
  }) */


  return NextResponse.json({});
}
