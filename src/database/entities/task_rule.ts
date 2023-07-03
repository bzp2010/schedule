import {
  Collection,
  Entity,
  Enum,
  ManyToOne,
  OneToMany,
  Property,
  type Rel,
} from "@mikro-orm/core";
import { BaseEntity, Status } from "@/database/entities/common";
import { Task } from "@/database/entities/task";
import { Job } from "@/database/entities/job";
import { TaskRule as GQLTaskRule } from "@/generated/graphql";

@Entity({ tableName: "task_rules" })
export class TaskRule extends BaseEntity {
  constructor(
    task: Rel<Task>,
    rule: string,
    description?: string,
    status?: Status
  ) {
    super();
    this.task = task;
    this.rule = rule;
    if (description) this.description = description;
    if (status) this.status = status;
  }

  @ManyToOne(() => Task)
  task: Rel<Task>;

  @Property()
  description: string = "";

  @Property()
  rule: string;

  @OneToMany(() => Job, (job) => job.taskRule)
  jobs = new Collection<Job>(this);

  @Property({ fieldName: "last_running_at", nullable: true })
  lastRunningAt!: Date;

  @Property({ fieldName: "last_running_time", nullable: true })
  lastRunningTime!: number;

  @Enum({ items: () => Status, type: "Status" })
  status: Status = Status.Enabled;

  toGQLType(): GQLTaskRule {
    return {
      id: this.id,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
      description: this.description,
      jobs: [],
      rule: this.rule,
      task: this.task && this.task.toGQLType(),
      status: this.status,
    };
  }
}
