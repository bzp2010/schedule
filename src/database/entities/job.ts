import { Entity, Enum, ManyToOne, Property, type Rel } from "@mikro-orm/core";
import { BaseEntity, Status } from "@/database/entities/common";
import { Task } from "@/database/entities/task";
import { TaskRule } from "@/database/entities/task_rule";
import { Job as GQLJob } from "@/generated/graphql";

export enum JobFlag {
  Timeout = 1,
}

@Entity({ tableName: "jobs" })
export class Job extends BaseEntity {
  constructor(
    task: Rel<Task>,
    taskRule: Rel<TaskRule>,
    startAt: Date,
    stopAt: Date,
    flags: number = 0
  ) {
    super();
    this.task = task;
    this.taskRule = taskRule;
    this.startAt = startAt;
    this.stopAt = stopAt;
    this.flags = flags;
  }

  @ManyToOne(() => Task)
  task: Rel<Task>;

  @ManyToOne(() => TaskRule)
  taskRule: Rel<TaskRule>;

  @Property()
  stdout: string = "";

  @Property()
  stderr: string = "";

  @Property()
  startAt: Date;

  @Property()
  stopAt: Date;

  @Property()
  flags: number = 0;

  @Property({ persist: false })
  get isTimeout() {
    return (this.flags & JobFlag.Timeout) != 0;
  }

  toGQLType(): GQLJob {
    return {
      id: this.id,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
      stdout: this.stdout,
      stderr: this.stderr,
      startAt: this.startAt,
      stopAt: this.stopAt,
      task: this.task && this.task.toGQLType(),
      taskRule: this.taskRule && this.taskRule.toGQLType(),
      isTimeout: this.isTimeout,
    };
  }
}
