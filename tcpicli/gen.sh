#!/bin/bash
############################################

# File Name : gen.sh

# Purpose :

# Creation Date : 11-02-2017

# Last Modified : Thu 02 Nov 2017 10:24:31 PM UTC

# Created By : Kiyor 

############################################

if [ -z $1 ]; then
	echo "Usage: $0 serviceName action1 action2 ..."
	exit
fi

for i in ${@:2}; do
	doc=`grep '// Implement' ../$1/${i}.go|awk '{ print $3 }'`
	echo "
		{
			Name:        \"$i\",
			Usage:       \"$i\",
			Action:      ${1^}$i,
			Description: \"referer $doc\",
		},
"
done

for i in ${@:2}; do
	echo "
func ${1^}$i(c *cli.Context) error {
	resp, err := $1.$i(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
"
done
