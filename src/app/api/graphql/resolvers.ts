import { Resolvers, Status, Job, Task, TaskRule } from "@/generated/graphql";
import { ApolloServerContext } from "./server";
import {
  Task as TaskEntity,
  TaskRule as TaskRuleEntity,
  Job as JobEntity,
} from "@/database/entities";
import { DateTimeResolver } from "graphql-scalars";

export const resolvers: Resolvers<ApolloServerContext> = {
  DateTime: DateTimeResolver,

  Query: {
    task: async (_, args, { em }) => {
      return (await em.findOneOrFail(TaskEntity, args.id)).toGQLType();
    },
    taskRule: async (_, args, { em }) => {
      return (await em.findOneOrFail(TaskRuleEntity, args.id)).toGQLType();
    },
    taskRules: async (_, args, { em }) => {
      return (
        await em.find(
          TaskRuleEntity,
          args.taskId
            ? {
                task: args.taskId,
              }
            : {},
          {
            limit: args.limit,
            offset: args.offset,
          }
        )
      ).map((rule) => rule.toGQLType());
    },
  },
  TaskConfiguration: {
    __resolveType: (obj) => {
      if ("command" in obj) return "TaskConfigurationShell";
      if ("method" in obj) return "TaskConfigurationWebhook";
      return null;
    },
  },
  Task: {
    rules: async (parent, args, { em }, info) => {
      return (
        await em.find(
          TaskRuleEntity,
          {
            task: parent.id,
          },
          {
            limit: args.limit,
            offset: args.offset,
          }
        )
      ).map((rule) => rule.toGQLType());
    },
    jobs: async (parent, args, { em }, info) => {
      return (
        await em.find(
          JobEntity,
          {
            task: parent.id,
          },
          {
            limit: args.limit,
            offset: args.offset,
          }
        )
      ).map((job) => job.toGQLType());
    },
    status: (parent) => {
      if (typeof parent.status == "number") {
        return parent.status === 1 ? Status.Enabled : Status.Disabled;
      }
      return parent.status;
    },
  },
  TaskRule: {
    description: (parent) => {
      return parent.description ?? "";
    },
    task: async (parent, _, { em }) => {
      return (
        await em.findOneOrFail(TaskEntity, {
          id: parent.task.id,
        })
      ).toGQLType();
    },
    jobs: async (parent, args, { em }) => {
      return (
        await em.find(
          JobEntity,
          {
            taskRule: parent.id,
          },
          {
            limit: args.limit,
            offset: args.offset,
          }
        )
      ).map((job) => job.toGQLType());
    },
    status: (parent) => {
      if (typeof parent.status == "number") {
        return parent.status === 1 ? Status.Enabled : Status.Disabled;
      }
      return parent.status;
    },
  },
};
