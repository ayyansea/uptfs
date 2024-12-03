# Undeclared Purpose Text Filtering System (uptfs)

This program takes some text, slices it to a number of tokens and applies filters to them.

uptfs will be considered complete as soon as all these statements are true:

* it can read text either from stdin or from a plain text file
* it can output text either to stdout or to a file
* it can be configured either fully with command line options or with a YAML file (or with a mix of both)
* it has the following configuration options (both in CLI and file modes):
    * -c, --config-file - path to the config file (DONE)
    * -i, --input-file - path to the input file
    * -o, --output-file - path to the output file
    * -f, --filter - name of a filter to apply (can be used multiple times) (DONE)
* it prioritizes command line options over those in the config file (DONE)
* it contains at least 7 different predefined filters (DONE)
* it has 'normal' and 'verbose' modes, the latter meaning presence of debug logs while the program is running (DONE)

Doing more doesn't really make sense since this is a learning project.
