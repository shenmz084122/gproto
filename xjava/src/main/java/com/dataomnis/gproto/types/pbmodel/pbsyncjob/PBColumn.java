// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/types/model/syncjob/column.proto

package com.dataomnis.gproto.types.pbmodel.pbsyncjob;

public final class PBColumn {
  private PBColumn() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ColumnOrBuilder extends
      // @@protoc_insertion_point(interface_extends:model.Column)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <pre>
     * column name
     * &#64;inject_tag: json:"name" 
     * </pre>
     *
     * <code>string name = 1;</code>
     * @return The name.
     */
    java.lang.String getName();
    /**
     * <pre>
     * column name
     * &#64;inject_tag: json:"name" 
     * </pre>
     *
     * <code>string name = 1;</code>
     * @return The bytes for name.
     */
    com.google.protobuf.ByteString
        getNameBytes();

    /**
     * <pre>
     * column type
     * &#64;inject_tag: json:"type" 
     * </pre>
     *
     * <code>string type = 2;</code>
     * @return The type.
     */
    java.lang.String getType();
    /**
     * <pre>
     * column type
     * &#64;inject_tag: json:"type" 
     * </pre>
     *
     * <code>string type = 2;</code>
     * @return The bytes for type.
     */
    com.google.protobuf.ByteString
        getTypeBytes();

    /**
     * <pre>
     * column index
     * &#64;inject_tag: json:"index"
     * </pre>
     *
     * <code>int32 index = 3;</code>
     * @return The index.
     */
    int getIndex();
  }
  /**
   * Protobuf type {@code model.Column}
   */
  public static final class Column extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:model.Column)
      ColumnOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use Column.newBuilder() to construct.
    private Column(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private Column() {
      name_ = "";
      type_ = "";
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new Column();
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private Column(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();

              name_ = s;
              break;
            }
            case 18: {
              java.lang.String s = input.readStringRequireUtf8();

              type_ = s;
              break;
            }
            case 24: {

              index_ = input.readInt32();
              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.internal_static_model_Column_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.internal_static_model_Column_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.class, com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.Builder.class);
    }

    /**
     * Protobuf enum {@code model.Column.ColumnMapping}
     */
    public enum ColumnMapping
        implements com.google.protobuf.ProtocolMessageEnum {
      /**
       * <code>Name = 0;</code>
       */
      Name(0),
      /**
       * <code>Row = 1;</code>
       */
      Row(1),
      /**
       * <code>Auto = 2;</code>
       */
      Auto(2),
      UNRECOGNIZED(-1),
      ;

      /**
       * <code>Name = 0;</code>
       */
      public static final int Name_VALUE = 0;
      /**
       * <code>Row = 1;</code>
       */
      public static final int Row_VALUE = 1;
      /**
       * <code>Auto = 2;</code>
       */
      public static final int Auto_VALUE = 2;


      public final int getNumber() {
        if (this == UNRECOGNIZED) {
          throw new java.lang.IllegalArgumentException(
              "Can't get the number of an unknown enum value.");
        }
        return value;
      }

      /**
       * @param value The numeric wire value of the corresponding enum entry.
       * @return The enum associated with the given numeric wire value.
       * @deprecated Use {@link #forNumber(int)} instead.
       */
      @java.lang.Deprecated
      public static ColumnMapping valueOf(int value) {
        return forNumber(value);
      }

      /**
       * @param value The numeric wire value of the corresponding enum entry.
       * @return The enum associated with the given numeric wire value.
       */
      public static ColumnMapping forNumber(int value) {
        switch (value) {
          case 0: return Name;
          case 1: return Row;
          case 2: return Auto;
          default: return null;
        }
      }

      public static com.google.protobuf.Internal.EnumLiteMap<ColumnMapping>
          internalGetValueMap() {
        return internalValueMap;
      }
      private static final com.google.protobuf.Internal.EnumLiteMap<
          ColumnMapping> internalValueMap =
            new com.google.protobuf.Internal.EnumLiteMap<ColumnMapping>() {
              public ColumnMapping findValueByNumber(int number) {
                return ColumnMapping.forNumber(number);
              }
            };

      public final com.google.protobuf.Descriptors.EnumValueDescriptor
          getValueDescriptor() {
        if (this == UNRECOGNIZED) {
          throw new java.lang.IllegalStateException(
              "Can't get the descriptor of an unrecognized enum value.");
        }
        return getDescriptor().getValues().get(ordinal());
      }
      public final com.google.protobuf.Descriptors.EnumDescriptor
          getDescriptorForType() {
        return getDescriptor();
      }
      public static final com.google.protobuf.Descriptors.EnumDescriptor
          getDescriptor() {
        return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.getDescriptor().getEnumTypes().get(0);
      }

      private static final ColumnMapping[] VALUES = values();

      public static ColumnMapping valueOf(
          com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
        if (desc.getType() != getDescriptor()) {
          throw new java.lang.IllegalArgumentException(
            "EnumValueDescriptor is not for this type.");
        }
        if (desc.getIndex() == -1) {
          return UNRECOGNIZED;
        }
        return VALUES[desc.getIndex()];
      }

      private final int value;

      private ColumnMapping(int value) {
        this.value = value;
      }

      // @@protoc_insertion_point(enum_scope:model.Column.ColumnMapping)
    }

    public static final int NAME_FIELD_NUMBER = 1;
    private volatile java.lang.Object name_;
    /**
     * <pre>
     * column name
     * &#64;inject_tag: json:"name" 
     * </pre>
     *
     * <code>string name = 1;</code>
     * @return The name.
     */
    @java.lang.Override
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      }
    }
    /**
     * <pre>
     * column name
     * &#64;inject_tag: json:"name" 
     * </pre>
     *
     * <code>string name = 1;</code>
     * @return The bytes for name.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int TYPE_FIELD_NUMBER = 2;
    private volatile java.lang.Object type_;
    /**
     * <pre>
     * column type
     * &#64;inject_tag: json:"type" 
     * </pre>
     *
     * <code>string type = 2;</code>
     * @return The type.
     */
    @java.lang.Override
    public java.lang.String getType() {
      java.lang.Object ref = type_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        type_ = s;
        return s;
      }
    }
    /**
     * <pre>
     * column type
     * &#64;inject_tag: json:"type" 
     * </pre>
     *
     * <code>string type = 2;</code>
     * @return The bytes for type.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getTypeBytes() {
      java.lang.Object ref = type_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        type_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int INDEX_FIELD_NUMBER = 3;
    private int index_;
    /**
     * <pre>
     * column index
     * &#64;inject_tag: json:"index"
     * </pre>
     *
     * <code>int32 index = 3;</code>
     * @return The index.
     */
    @java.lang.Override
    public int getIndex() {
      return index_;
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(name_)) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, name_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(type_)) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 2, type_);
      }
      if (index_ != 0) {
        output.writeInt32(3, index_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(name_)) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, name_);
      }
      if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(type_)) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, type_);
      }
      if (index_ != 0) {
        size += com.google.protobuf.CodedOutputStream
          .computeInt32Size(3, index_);
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column)) {
        return super.equals(obj);
      }
      com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column other = (com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column) obj;

      if (!getName()
          .equals(other.getName())) return false;
      if (!getType()
          .equals(other.getType())) return false;
      if (getIndex()
          != other.getIndex()) return false;
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + NAME_FIELD_NUMBER;
      hash = (53 * hash) + getName().hashCode();
      hash = (37 * hash) + TYPE_FIELD_NUMBER;
      hash = (53 * hash) + getType().hashCode();
      hash = (37 * hash) + INDEX_FIELD_NUMBER;
      hash = (53 * hash) + getIndex();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code model.Column}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:model.Column)
        com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.ColumnOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.internal_static_model_Column_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.internal_static_model_Column_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.class, com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.Builder.class);
      }

      // Construct using com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        name_ = "";

        type_ = "";

        index_ = 0;

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.internal_static_model_Column_descriptor;
      }

      @java.lang.Override
      public com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column getDefaultInstanceForType() {
        return com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.getDefaultInstance();
      }

      @java.lang.Override
      public com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column build() {
        com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column buildPartial() {
        com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column result = new com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column(this);
        result.name_ = name_;
        result.type_ = type_;
        result.index_ = index_;
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column) {
          return mergeFrom((com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column other) {
        if (other == com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column.getDefaultInstance()) return this;
        if (!other.getName().isEmpty()) {
          name_ = other.name_;
          onChanged();
        }
        if (!other.getType().isEmpty()) {
          type_ = other.type_;
          onChanged();
        }
        if (other.getIndex() != 0) {
          setIndex(other.getIndex());
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private java.lang.Object name_ = "";
      /**
       * <pre>
       * column name
       * &#64;inject_tag: json:"name" 
       * </pre>
       *
       * <code>string name = 1;</code>
       * @return The name.
       */
      public java.lang.String getName() {
        java.lang.Object ref = name_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          name_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <pre>
       * column name
       * &#64;inject_tag: json:"name" 
       * </pre>
       *
       * <code>string name = 1;</code>
       * @return The bytes for name.
       */
      public com.google.protobuf.ByteString
          getNameBytes() {
        java.lang.Object ref = name_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          name_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <pre>
       * column name
       * &#64;inject_tag: json:"name" 
       * </pre>
       *
       * <code>string name = 1;</code>
       * @param value The name to set.
       * @return This builder for chaining.
       */
      public Builder setName(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        name_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * column name
       * &#64;inject_tag: json:"name" 
       * </pre>
       *
       * <code>string name = 1;</code>
       * @return This builder for chaining.
       */
      public Builder clearName() {
        
        name_ = getDefaultInstance().getName();
        onChanged();
        return this;
      }
      /**
       * <pre>
       * column name
       * &#64;inject_tag: json:"name" 
       * </pre>
       *
       * <code>string name = 1;</code>
       * @param value The bytes for name to set.
       * @return This builder for chaining.
       */
      public Builder setNameBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        name_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object type_ = "";
      /**
       * <pre>
       * column type
       * &#64;inject_tag: json:"type" 
       * </pre>
       *
       * <code>string type = 2;</code>
       * @return The type.
       */
      public java.lang.String getType() {
        java.lang.Object ref = type_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          type_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <pre>
       * column type
       * &#64;inject_tag: json:"type" 
       * </pre>
       *
       * <code>string type = 2;</code>
       * @return The bytes for type.
       */
      public com.google.protobuf.ByteString
          getTypeBytes() {
        java.lang.Object ref = type_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          type_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <pre>
       * column type
       * &#64;inject_tag: json:"type" 
       * </pre>
       *
       * <code>string type = 2;</code>
       * @param value The type to set.
       * @return This builder for chaining.
       */
      public Builder setType(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        type_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * column type
       * &#64;inject_tag: json:"type" 
       * </pre>
       *
       * <code>string type = 2;</code>
       * @return This builder for chaining.
       */
      public Builder clearType() {
        
        type_ = getDefaultInstance().getType();
        onChanged();
        return this;
      }
      /**
       * <pre>
       * column type
       * &#64;inject_tag: json:"type" 
       * </pre>
       *
       * <code>string type = 2;</code>
       * @param value The bytes for type to set.
       * @return This builder for chaining.
       */
      public Builder setTypeBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        type_ = value;
        onChanged();
        return this;
      }

      private int index_ ;
      /**
       * <pre>
       * column index
       * &#64;inject_tag: json:"index"
       * </pre>
       *
       * <code>int32 index = 3;</code>
       * @return The index.
       */
      @java.lang.Override
      public int getIndex() {
        return index_;
      }
      /**
       * <pre>
       * column index
       * &#64;inject_tag: json:"index"
       * </pre>
       *
       * <code>int32 index = 3;</code>
       * @param value The index to set.
       * @return This builder for chaining.
       */
      public Builder setIndex(int value) {
        
        index_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * column index
       * &#64;inject_tag: json:"index"
       * </pre>
       *
       * <code>int32 index = 3;</code>
       * @return This builder for chaining.
       */
      public Builder clearIndex() {
        
        index_ = 0;
        onChanged();
        return this;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:model.Column)
    }

    // @@protoc_insertion_point(class_scope:model.Column)
    private static final com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column();
    }

    public static com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Column>
        PARSER = new com.google.protobuf.AbstractParser<Column>() {
      @java.lang.Override
      public Column parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new Column(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<Column> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Column> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public com.dataomnis.gproto.types.pbmodel.pbsyncjob.PBColumn.Column getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_model_Column_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_model_Column_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n&proto/types/model/syncjob/column.proto" +
      "\022\005model\0323github.com/yu31/protoc-plugin/p" +
      "roto/validator.proto\032/github.com/yu31/pr" +
      "otoc-plugin/proto/gosql.proto\0322github.co" +
      "m/yu31/protoc-plugin/proto/defaults.prot" +
      "o\"a\n\006Column\022\014\n\004name\030\001 \001(\t\022\014\n\004type\030\002 \001(\t\022" +
      "\r\n\005index\030\003 \001(\005\",\n\rColumnMapping\022\010\n\004Name\020" +
      "\000\022\007\n\003Row\020\001\022\010\n\004Auto\020\002Bw\n,com.dataomnis.gp" +
      "roto.types.pbmodel.pbsyncjobB\010PBColumnP\000" +
      "Z;github.com/DataWorkbench/gproto/xgo/ty" +
      "pes/pbmodel/pbsyncjobb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.github.yu31.protoc.pb.pbvalidator.PBValidator.getDescriptor(),
          io.github.yu31.protoc.pb.pbgosql.PBGoSQL.getDescriptor(),
          io.github.yu31.protoc.pb.pbdefaults.PBDefaults.getDescriptor(),
        });
    internal_static_model_Column_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_model_Column_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_model_Column_descriptor,
        new java.lang.String[] { "Name", "Type", "Index", });
    io.github.yu31.protoc.pb.pbvalidator.PBValidator.getDescriptor();
    io.github.yu31.protoc.pb.pbgosql.PBGoSQL.getDescriptor();
    io.github.yu31.protoc.pb.pbdefaults.PBDefaults.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
