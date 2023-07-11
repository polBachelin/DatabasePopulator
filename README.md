# Database Populator

DatabasePopulator is a command line tool that helps developers
populate any mongoDB with as many documents as they want.

This application is a tool to generate large quantities of documents for testing tools on local databases

## Commands

- generate

## Usage/Examples

To use this CLI you need to provide some flags

### Database

The database flag with the necessary information for the tool to connect to the database

```
    -host localhost
    -user root
    -pass 12345
    -port 27017
```

### Schema

The tool uses _YAML_ files to define the document structure of your database.
You can specify a folder or only a file to be used for the document generation with the flag `-schema`

You can find examples of document schemas in the [examples](https://github.com/polBachelin/DatabasePopulator/tree/main/examples) folder

A collection in this tool is represented by a block. The name of the block is what will be used to create the collection. The documents created in this collection are defined by the fields that you specify in your _YAML_ file

The generated documents are comprised of fields. These fields have three components

```
    - name: <fieldName>
      type: <fieldType>
      default_value: <defaultValue>
```

a field can have any of the following types

- string
- date
- integer
- double
- boolean
- array
- object
- timestamp
- binary

#### Default Value

A default value is optional.

This default value is only used if the flag -default is set on the relative fields. To mention a specific field you need to follow the structure [BLOCK_NAME.FIELD_NAME]

For example if I want the generated documents to use the default value of the field **category** in the collection **Stories**. I need to set `-default Stories.category`. If no default_value is set in the YAML file an error is thrown

---

#### Object

An object is a nested document. This is represented by a simply nesting the fields after the type

```
- name: movies
  type: object
  fields:
    - name: release_date
      type: date
```

If no fields is specified after the object type is used an error is thrown

### Generation

The document generation can be customized to your needs. Here are the following flags that can be used

```
-count <block_name> 100 : the number of a specific document you want to generate.
If no block_name is specified then the count is used for every collection

-value <block_name.field_name> <value> <value>... : the value to use for the generation
Any number of values can be added and the generator will randomly choose one
If no value is specified for a field then a random value will be generated

-value_beetween <block_name.field_name> <start> <end> : a start and end value for the generator to create values in beetween. Mostly used for date types
```
