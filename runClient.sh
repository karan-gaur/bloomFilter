#!/bin/bash

data="sample_docs"
file_search="false"
keywords="keywords.txt"
num_docs="128"
bf_sz="1024"
hash_lvl="5"

while getopts ":h?:d:t:b:n:m:f:s:c:x:y:q:r:p:z:l:a:u:" opt; do
    case "$opt" in
        h|\?)
            echo -e "\nArguments: "
            echo -e "-b \t\t Bits in Bloom filter (default 1024)"
            echo -e "-n \t\t Max number of documents (default 128)"
            echo -e "-f \t\t File search - will search keywords from file (default 'false')"
            echo -e "-k \t\t keywords filename to search if 'File Search' is set to 'true' (default - 'keywords.txt')"
            echo -e "-s \t\t Name of input directory for generating Bloom Filter (default - 'sample_docs'\n"
            echo -e "-l \t\t Hash Level for the Bloom Filter Hash function (default '5')"
            ;;
        b)
            bf_sz=$OPTARG
            ;;
        n)
            num_docs=$OPTARG
            ;;
        f)
            file_search=$OPTARG
            ;;
        k)
            keywords=$OPTARG
            ;;
        s)
            data=$OPTARG
            ;;
        l)
            hash_lvl=$OPTARG
            ;;
    esac
done

echo "data='$data', file_search='$file_search', keywords='$keywords', bf_sz='$bf_sz', num_docs='$num_docs', hash_lvl='$hash_lvl'"

go run src/bench/main.go --data="$data" --file_search="$file_search" --keywords="$keywords" --num_docs="$num_docs" --bf_sz="$bf_sz" --hash_lvl="$hash_lvl"