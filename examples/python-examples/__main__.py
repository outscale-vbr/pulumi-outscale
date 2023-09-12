"""A Python Pulumi program"""

import pulumi
import pulumi_outscale as outscale 

deployer = outscale.Keypair("deployer", keypair_name="deployer")