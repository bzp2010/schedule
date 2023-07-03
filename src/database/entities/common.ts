import { PrimaryKey, Property } from "@mikro-orm/core";

export abstract class BaseEntity {
  @PrimaryKey({ fieldName: "id" })
  readonly id!: number;

  @Property({ fieldName: "created_at" })
  readonly createdAt: Date = new Date();

  @Property({ fieldName: "updated_at", onUpdate: () => new Date() })
  readonly updatedAt: Date = new Date();

  @Property({ fieldName: "deleted_at", nullable: true })
  readonly deletedAt?: Date | undefined = undefined;

  abstract toGQLType(): any;
}

export enum Status {
  Disabled,
  Enabled,
}
