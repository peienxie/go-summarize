# Summarization CLI in golang

This is a simple CLI application built in Go that uses the OpenAI API to summarize input text.

It can accept input from either a file or a string.

# Prerequisites

Before running this application, you will need to set up an OpenAI API key.

You can get a key by creating an account at [OpenAI](https://platform.openai.com/account/api-keys.).

# How to Build

To install this application, clone this repository and build it by `go build`

# Usage

To run the application, set the `OPENAI_API_KEY` environment variable to
your OpenAI API key, and then run the executable:

```
$ export OPENAI_API_KEY=sk-BrpuQnoswqkq01230MUW8asdfwklefjlbkb
$ ./go-summarize --text "input you text here"`
```

## Summarize input text

`./go-summarize --text "input you text here"`

## Summarize text from a file

`./go-summarize --file "input you text here"`
