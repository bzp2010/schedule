import { Collection, Entity, Enum, OneToMany, Property } from "@mikro-orm/core";
import { BaseEntity, Status } from "@/database/entities/common";
import { TaskRule } from "@/database/entities/task_rule";
import { Job } from "@/database/entities/job";
import {
  Task as GQLTask,
  HttpMethod as GQLHttpMethod,
} from "@/generated/graphql";

export enum TaskType {
  Shell = "SHELL",
  Webhook = "WEBHOOK",
}

export interface TaskConfigurationShell {
  command: string;
  timeout: number;
}

export interface TaskConfigurationWebhook {
  url: string;
  method: GQLHttpMethod;
}

export type TaskConfiguration =
  | TaskConfigurationShell
  | TaskConfigurationWebhook;

@Entity({ tableName: "tasks" })
export class Task extends BaseEntity {
  constructor(
    name: string,
    type: TaskType,
    configuration: TaskConfiguration,
    status?: Status
  ) {
    super();
    this.name = name;
    this.type = type;
    this.configuration = configuration;
    if (status) this.status = status;
  }

  @Property()
  name: string;

  @Enum(() => TaskType)
  type: TaskType;

  @Property({ type: "json" })
  configuration: TaskConfiguration;

  @OneToMany(() => TaskRule, (rule) => rule.task)
  rules = new Collection<TaskRule>(this);

  @OneToMany(() => Job, (job) => job.task)
  jobs = new Collection<Job>(this);

  @Property({ fieldName: "last_running_at", nullable: true })
  lastRunningAt?: Date;

  @Property({ fieldName: "last_running_time", nullable: true })
  lastRunningTime?: number;

  @Enum({ items: () => Status, type: "Status" })
  status: Status = Status.Enabled;

  toGQLType(): GQLTask {
    return {
      id: this.id,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
      name: this.name,
      type: this.type,
      configuration: this.configuration,
      lastRunningAt: this.lastRunningAt,
      lastRunningTime: this.lastRunningTime,
      jobs: [],
      rules: [],
      status: this.status,
    };
  }
}
