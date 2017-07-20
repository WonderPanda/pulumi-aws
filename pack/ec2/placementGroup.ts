// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class PlacementGroup extends lumi.NamedResource implements PlacementGroupArgs {
    public readonly placementGroupName?: string;
    public readonly strategy: string;

    constructor(name: string, args: PlacementGroupArgs) {
        super(name);
        this.placementGroupName = args.placementGroupName;
        if (lumirt.defaultIfComputed(args.strategy, "") === undefined) {
            throw new Error("Property argument 'strategy' is required, but was missing");
        }
        this.strategy = args.strategy;
    }
}

export interface PlacementGroupArgs {
    readonly placementGroupName?: string;
    readonly strategy: string;
}
