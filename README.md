# gopt

## Goal

Goal of this project is to learn Go language.
Domain of the project will be JIRA-like system.
System will contain 'Projects'.
Each Project can  contain 0+ number of Items.
Items can be of different types (TBD).
It will be possible to define different lifecycles for different types of items.(?)
Different types of items can have  0+ Items as children.
There will be different  types of relations between Items on the same level, or parent-child relation.
HTML/HTMX/TEMPL frontend.

## Run
process TEMPL templates with templ generate
create postgres database
have DB_URL env variable, eg.: DB_URL=postgresql://postgres:postgres@localhost:5432/gopt
run app with -initdb true parameter

