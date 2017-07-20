// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class SqlInjectionMatchSet extends lumi.NamedResource implements SqlInjectionMatchSetArgs {
    public readonly sqlInjectionMatchSetName?: string;
    public readonly sqlInjectionMatchTuples?: { fieldToMatch: { data?: string, type: string }[], textTransformation: string }[];

    constructor(name: string, args: SqlInjectionMatchSetArgs) {
        super(name);
        this.sqlInjectionMatchSetName = args.sqlInjectionMatchSetName;
        this.sqlInjectionMatchTuples = args.sqlInjectionMatchTuples;
    }
}

export interface SqlInjectionMatchSetArgs {
    readonly sqlInjectionMatchSetName?: string;
    readonly sqlInjectionMatchTuples?: { fieldToMatch: { data?: string, type: string }[], textTransformation: string }[];
}
