import { DataFrame, DataFrameView } from '@grafana/data';

export interface GqlNode {
    gql: string;
    displayName: string;
    type: string;
    geoPath?: string;
}

export class FindMetricHelper {
    static toGqlNodeList(frame: DataFrame): GqlNode[] {
        const rows = new DataFrameView<GqlNode>(frame);
        return rows.toArray();
    }
}
