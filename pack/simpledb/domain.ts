// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Domain extends lumi.NamedResource implements DomainArgs {
    public readonly domainName?: string;

    constructor(name: string, args: DomainArgs) {
        super(name);
        this.domainName = args.domainName;
    }
}

export interface DomainArgs {
    readonly domainName?: string;
}
