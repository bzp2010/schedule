import { GraphQLResolveInfo, GraphQLScalarType, GraphQLScalarTypeConfig } from 'graphql';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  DateTime: { input: any; output: any; }
  /** Long integer scalar type definition */
  Int64: { input: any; output: any; }
};

/** CreateTask is the data structure used by mutation of create Task */
export type CreateTask = {
  /** Task configuration */
  configuration: InputTaskConfiguration;
  /** Task name */
  name: Scalars['String']['input'];
  /** Status of the task */
  status?: InputMaybe<Status>;
  /** Task type (SHELL, WEBHOOK) */
  type: TaskType;
};

/** EditTask is the data structure used by mutation of edit Task */
export type EditTask = {
  /** Task configuration */
  configuration?: InputMaybe<InputTaskConfiguration>;
  /** Task name */
  name?: InputMaybe<Scalars['String']['input']>;
  /** Status of the task */
  status?: InputMaybe<Status>;
  /** Task type (SHELL, WEBHOOK) */
  type?: InputMaybe<TaskType>;
};

/** HTTPMethod indicates the list of HTTP methods */
export enum HttpMethod {
  Connect = 'CONNECT',
  Delete = 'DELETE',
  Get = 'GET',
  Head = 'HEAD',
  Options = 'OPTIONS',
  Patch = 'PATCH',
  Post = 'POST',
  Put = 'PUT',
  Trace = 'TRACE'
}

/**
 * InputTaskConfiguration is a collection type of InputTaskConfigurationShell and InputTaskConfigurationWebhook,
 * which is an alternative to the temporarily unsupported inputUnion
 */
export type InputTaskConfiguration = {
  /** Shell task configuration */
  shell?: InputMaybe<InputTaskConfigurationShell>;
  /** Webhook task configuration */
  webhook?: InputMaybe<InputTaskConfigurationWebhook>;
};

/** InputTaskConfigurationShell is the configuration input for shell type tasks */
export type InputTaskConfigurationShell = {
  /** Command run in the Shell task */
  command: Scalars['String']['input'];
  /** Maximum time of Shell task execution */
  timeout: Scalars['Int64']['input'];
};

/** InputTaskConfigurationWebhook is the configuration input for webhook type tasks */
export type InputTaskConfigurationWebhook = {
  /** HTTP methods of the Webhook task request */
  method: HttpMethod;
  /** URL of the Webhook task request */
  url: Scalars['String']['input'];
};

export type Job = Model & {
  __typename?: 'Job';
  /** Create time of the job */
  createdAt: Scalars['DateTime']['output'];
  /** Job ID */
  id: Scalars['Int']['output'];
  /** Job execution timeout */
  isTimeout: Scalars['Boolean']['output'];
  /** Job execution starts at */
  startAt: Scalars['DateTime']['output'];
  /** The stderr of job */
  stderr: Scalars['String']['output'];
  /** The stdout of job */
  stdout: Scalars['String']['output'];
  /** Job execution stops at */
  stopAt: Scalars['DateTime']['output'];
  /** Job associated task */
  task: Task;
  /** Job associated task rule */
  taskRule: TaskRule;
  /** Last update time of the job */
  updatedAt: Scalars['DateTime']['output'];
};

/** Definition of the base fields of the data model */
export type Model = {
  /** Entry created time */
  createdAt: Scalars['DateTime']['output'];
  /** Entry ID */
  id: Scalars['Int']['output'];
  /** Entry updated time */
  updatedAt: Scalars['DateTime']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  /** Create a task */
  createTask: Task;
  /** Delete a task */
  deleteTask?: Maybe<Scalars['Boolean']['output']>;
  /** Edit a task */
  editTask?: Maybe<Task>;
};


export type MutationCreateTaskArgs = {
  input: CreateTask;
};


export type MutationDeleteTaskArgs = {
  id: Scalars['ID']['input'];
};


export type MutationEditTaskArgs = {
  id: Scalars['ID']['input'];
  input: EditTask;
};

export type Query = {
  __typename?: 'Query';
  /** Get a single job */
  job?: Maybe<Job>;
  /** Get the list of jobs */
  jobs: Array<Job>;
  /** Get a single task */
  task?: Maybe<Task>;
  /** Get a single task rule */
  taskRule?: Maybe<TaskRule>;
  /** Get the list of task rules */
  taskRules: Array<TaskRule>;
  /** Get the list of tasks */
  tasks: Array<Task>;
};


export type QueryJobArgs = {
  id: Scalars['Int']['input'];
};


export type QueryJobsArgs = {
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
  reverse_order?: Scalars['Boolean']['input'];
};


export type QueryTaskArgs = {
  id: Scalars['Int']['input'];
};


export type QueryTaskRuleArgs = {
  id: Scalars['Int']['input'];
};


export type QueryTaskRulesArgs = {
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
  taskId?: InputMaybe<Scalars['Int']['input']>;
};


export type QueryTasksArgs = {
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
};

/** Definition of the status of a data entry */
export enum Status {
  Disabled = 'DISABLED',
  Enabled = 'ENABLED'
}

/** Definition of Task data model */
export type Task = Model & {
  __typename?: 'Task';
  /** Task configuration */
  configuration: TaskConfiguration;
  /** Create time of the task */
  createdAt: Scalars['DateTime']['output'];
  /** Task ID */
  id: Scalars['Int']['output'];
  /** Jobs associated with the task */
  jobs: Array<Job>;
  /** Last execution moment of the task */
  lastRunningAt?: Maybe<Scalars['DateTime']['output']>;
  /** Last execution time of the task */
  lastRunningTime?: Maybe<Scalars['Int']['output']>;
  /** Task name */
  name: Scalars['String']['output'];
  /** Rules of the task */
  rules: Array<TaskRule>;
  /** Status of the task */
  status: Status;
  /** Task type (SHELL, WEBHOOK) */
  type: TaskType;
  /** Last update time of the task */
  updatedAt: Scalars['DateTime']['output'];
};


/** Definition of Task data model */
export type TaskJobsArgs = {
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
  reverse_order?: Scalars['Boolean']['input'];
};


/** Definition of Task data model */
export type TaskRulesArgs = {
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
};

/** TaskConfiguration is an aggregated type of multiple configurations */
export type TaskConfiguration = TaskConfigurationShell | TaskConfigurationWebhook;

/** TaskConfigurationShell is the configuration for shell type tasks */
export type TaskConfigurationShell = {
  __typename?: 'TaskConfigurationShell';
  /** Command run in the Shell task */
  command: Scalars['String']['output'];
  /** Maximum time of Shell task execution */
  timeout: Scalars['Int64']['output'];
};

/** TaskConfigurationShell is the configuration for webhook type tasks */
export type TaskConfigurationWebhook = {
  __typename?: 'TaskConfigurationWebhook';
  /** HTTP methods of the Webhook task request */
  method: HttpMethod;
  /** URL of the Webhook task request */
  url: Scalars['String']['output'];
};

export type TaskRule = Model & {
  __typename?: 'TaskRule';
  /** Create time of the task */
  createdAt: Scalars['DateTime']['output'];
  /** Task rule description */
  description: Scalars['String']['output'];
  /** Task rule ID */
  id: Scalars['Int']['output'];
  /** Jobs associated with the rule */
  jobs: Array<Job>;
  /** Last execution moment of the task */
  lastRunningAt?: Maybe<Scalars['DateTime']['output']>;
  /** Last execution time of the task */
  lastRunningTime?: Maybe<Scalars['Int']['output']>;
  /** Task rule definition (in quartz format) */
  rule: Scalars['String']['output'];
  /** Status of the task */
  status: Status;
  /** Task associated with the rule */
  task: Task;
  /** Last update time of the task */
  updatedAt: Scalars['DateTime']['output'];
};


export type TaskRuleJobsArgs = {
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
  reverse_order?: Scalars['Boolean']['input'];
};

/** TaskType indicates the list of task type */
export enum TaskType {
  /** The type of task that execute external shell commands and log the results */
  Shell = 'SHELL',
  /** The type of task that requests an external webhook and logs the results */
  Webhook = 'WEBHOOK'
}



export type ResolverTypeWrapper<T> = Promise<T> | T;


export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};
export type Resolver<TResult, TParent = {}, TContext = {}, TArgs = {}> = ResolverFn<TResult, TParent, TContext, TArgs> | ResolverWithResolve<TResult, TParent, TContext, TArgs>;

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => Promise<TResult> | TResult;

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>;

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>;
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>;
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>;
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>;
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>;

export type SubscriptionResolver<TResult, TKey extends string, TParent = {}, TContext = {}, TArgs = {}> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>;

export type TypeResolveFn<TTypes, TParent = {}, TContext = {}> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo
) => Maybe<TTypes> | Promise<Maybe<TTypes>>;

export type IsTypeOfResolverFn<T = {}, TContext = {}> = (obj: T, context: TContext, info: GraphQLResolveInfo) => boolean | Promise<boolean>;

export type NextResolverFn<T> = () => Promise<T>;

export type DirectiveResolverFn<TResult = {}, TParent = {}, TContext = {}, TArgs = {}> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

/** Mapping of union types */
export type ResolversUnionTypes<RefType extends Record<string, unknown>> = {
  TaskConfiguration: ( TaskConfigurationShell ) | ( TaskConfigurationWebhook );
};

/** Mapping of interface types */
export type ResolversInterfaceTypes<RefType extends Record<string, unknown>> = {
  Model: ( Job ) | ( Omit<Task, 'configuration'> & { configuration: RefType['TaskConfiguration'] } ) | ( TaskRule );
};

/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = {
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  CreateTask: CreateTask;
  DateTime: ResolverTypeWrapper<Scalars['DateTime']['output']>;
  EditTask: EditTask;
  HTTPMethod: HttpMethod;
  ID: ResolverTypeWrapper<Scalars['ID']['output']>;
  InputTaskConfiguration: InputTaskConfiguration;
  InputTaskConfigurationShell: InputTaskConfigurationShell;
  InputTaskConfigurationWebhook: InputTaskConfigurationWebhook;
  Int: ResolverTypeWrapper<Scalars['Int']['output']>;
  Int64: ResolverTypeWrapper<Scalars['Int64']['output']>;
  Job: ResolverTypeWrapper<Job>;
  Model: ResolverTypeWrapper<ResolversInterfaceTypes<ResolversTypes>['Model']>;
  Mutation: ResolverTypeWrapper<{}>;
  Query: ResolverTypeWrapper<{}>;
  Status: Status;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  Task: ResolverTypeWrapper<Omit<Task, 'configuration'> & { configuration: ResolversTypes['TaskConfiguration'] }>;
  TaskConfiguration: ResolverTypeWrapper<ResolversUnionTypes<ResolversTypes>['TaskConfiguration']>;
  TaskConfigurationShell: ResolverTypeWrapper<TaskConfigurationShell>;
  TaskConfigurationWebhook: ResolverTypeWrapper<TaskConfigurationWebhook>;
  TaskRule: ResolverTypeWrapper<TaskRule>;
  TaskType: TaskType;
};

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = {
  Boolean: Scalars['Boolean']['output'];
  CreateTask: CreateTask;
  DateTime: Scalars['DateTime']['output'];
  EditTask: EditTask;
  ID: Scalars['ID']['output'];
  InputTaskConfiguration: InputTaskConfiguration;
  InputTaskConfigurationShell: InputTaskConfigurationShell;
  InputTaskConfigurationWebhook: InputTaskConfigurationWebhook;
  Int: Scalars['Int']['output'];
  Int64: Scalars['Int64']['output'];
  Job: Job;
  Model: ResolversInterfaceTypes<ResolversParentTypes>['Model'];
  Mutation: {};
  Query: {};
  String: Scalars['String']['output'];
  Task: Omit<Task, 'configuration'> & { configuration: ResolversParentTypes['TaskConfiguration'] };
  TaskConfiguration: ResolversUnionTypes<ResolversParentTypes>['TaskConfiguration'];
  TaskConfigurationShell: TaskConfigurationShell;
  TaskConfigurationWebhook: TaskConfigurationWebhook;
  TaskRule: TaskRule;
};

export interface DateTimeScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['DateTime'], any> {
  name: 'DateTime';
}

export interface Int64ScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['Int64'], any> {
  name: 'Int64';
}

export type JobResolvers<ContextType = any, ParentType extends ResolversParentTypes['Job'] = ResolversParentTypes['Job']> = {
  createdAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  isTimeout?: Resolver<ResolversTypes['Boolean'], ParentType, ContextType>;
  startAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  stderr?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  stdout?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  stopAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  task?: Resolver<ResolversTypes['Task'], ParentType, ContextType>;
  taskRule?: Resolver<ResolversTypes['TaskRule'], ParentType, ContextType>;
  updatedAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type ModelResolvers<ContextType = any, ParentType extends ResolversParentTypes['Model'] = ResolversParentTypes['Model']> = {
  __resolveType: TypeResolveFn<'Job' | 'Task' | 'TaskRule', ParentType, ContextType>;
  createdAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  updatedAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
};

export type MutationResolvers<ContextType = any, ParentType extends ResolversParentTypes['Mutation'] = ResolversParentTypes['Mutation']> = {
  createTask?: Resolver<ResolversTypes['Task'], ParentType, ContextType, RequireFields<MutationCreateTaskArgs, 'input'>>;
  deleteTask?: Resolver<Maybe<ResolversTypes['Boolean']>, ParentType, ContextType, RequireFields<MutationDeleteTaskArgs, 'id'>>;
  editTask?: Resolver<Maybe<ResolversTypes['Task']>, ParentType, ContextType, RequireFields<MutationEditTaskArgs, 'id' | 'input'>>;
};

export type QueryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = {
  job?: Resolver<Maybe<ResolversTypes['Job']>, ParentType, ContextType, RequireFields<QueryJobArgs, 'id'>>;
  jobs?: Resolver<Array<ResolversTypes['Job']>, ParentType, ContextType, RequireFields<QueryJobsArgs, 'limit' | 'offset' | 'reverse_order'>>;
  task?: Resolver<Maybe<ResolversTypes['Task']>, ParentType, ContextType, RequireFields<QueryTaskArgs, 'id'>>;
  taskRule?: Resolver<Maybe<ResolversTypes['TaskRule']>, ParentType, ContextType, RequireFields<QueryTaskRuleArgs, 'id'>>;
  taskRules?: Resolver<Array<ResolversTypes['TaskRule']>, ParentType, ContextType, RequireFields<QueryTaskRulesArgs, 'limit' | 'offset'>>;
  tasks?: Resolver<Array<ResolversTypes['Task']>, ParentType, ContextType, RequireFields<QueryTasksArgs, 'limit' | 'offset'>>;
};

export type TaskResolvers<ContextType = any, ParentType extends ResolversParentTypes['Task'] = ResolversParentTypes['Task']> = {
  configuration?: Resolver<ResolversTypes['TaskConfiguration'], ParentType, ContextType>;
  createdAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  jobs?: Resolver<Array<ResolversTypes['Job']>, ParentType, ContextType, RequireFields<TaskJobsArgs, 'limit' | 'offset' | 'reverse_order'>>;
  lastRunningAt?: Resolver<Maybe<ResolversTypes['DateTime']>, ParentType, ContextType>;
  lastRunningTime?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  rules?: Resolver<Array<ResolversTypes['TaskRule']>, ParentType, ContextType, RequireFields<TaskRulesArgs, 'limit' | 'offset'>>;
  status?: Resolver<ResolversTypes['Status'], ParentType, ContextType>;
  type?: Resolver<ResolversTypes['TaskType'], ParentType, ContextType>;
  updatedAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type TaskConfigurationResolvers<ContextType = any, ParentType extends ResolversParentTypes['TaskConfiguration'] = ResolversParentTypes['TaskConfiguration']> = {
  __resolveType: TypeResolveFn<'TaskConfigurationShell' | 'TaskConfigurationWebhook', ParentType, ContextType>;
};

export type TaskConfigurationShellResolvers<ContextType = any, ParentType extends ResolversParentTypes['TaskConfigurationShell'] = ResolversParentTypes['TaskConfigurationShell']> = {
  command?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  timeout?: Resolver<ResolversTypes['Int64'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type TaskConfigurationWebhookResolvers<ContextType = any, ParentType extends ResolversParentTypes['TaskConfigurationWebhook'] = ResolversParentTypes['TaskConfigurationWebhook']> = {
  method?: Resolver<ResolversTypes['HTTPMethod'], ParentType, ContextType>;
  url?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type TaskRuleResolvers<ContextType = any, ParentType extends ResolversParentTypes['TaskRule'] = ResolversParentTypes['TaskRule']> = {
  createdAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  description?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  jobs?: Resolver<Array<ResolversTypes['Job']>, ParentType, ContextType, RequireFields<TaskRuleJobsArgs, 'limit' | 'offset' | 'reverse_order'>>;
  lastRunningAt?: Resolver<Maybe<ResolversTypes['DateTime']>, ParentType, ContextType>;
  lastRunningTime?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  rule?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  status?: Resolver<ResolversTypes['Status'], ParentType, ContextType>;
  task?: Resolver<ResolversTypes['Task'], ParentType, ContextType>;
  updatedAt?: Resolver<ResolversTypes['DateTime'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type Resolvers<ContextType = any> = {
  DateTime?: GraphQLScalarType;
  Int64?: GraphQLScalarType;
  Job?: JobResolvers<ContextType>;
  Model?: ModelResolvers<ContextType>;
  Mutation?: MutationResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
  Task?: TaskResolvers<ContextType>;
  TaskConfiguration?: TaskConfigurationResolvers<ContextType>;
  TaskConfigurationShell?: TaskConfigurationShellResolvers<ContextType>;
  TaskConfigurationWebhook?: TaskConfigurationWebhookResolvers<ContextType>;
  TaskRule?: TaskRuleResolvers<ContextType>;
};

