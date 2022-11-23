#bin/bash

AMI=ami-076309742d466ad69
sg=sg-09c40474f50fffa18

aws ec2 run-instances --image-id ${AMI} --count 1 \
--instance-type t2.micro \
--security-group-ids ${sg} \
--key-name key-sdcc --associate-public-ip-address \
--tag-specifications "ResourceType=instance,Tags=[{Key=Name,Value=SDCC-$1}]"