import {DataQueryResponse, DataSourceJsonData} from '@grafana/data';
import { DataQuery } from '@grafana/schema';
import {GqlNode} from "findMetricHelper";
import {AlcDataSource} from "datasource";

export interface AlcQuerySpec {
  alias: string;
  gql?: string;
}

/*
 * AlcQuery defines the format of the query specification sent to
 * the backend for each query (request).  Every query uses this
 * request format, regardless of the type of query
 */
export interface AlcQueryProps extends AlcQuerySpec, DataQuery {}

/*
 * FindChildrenRequest and FindChildrenResponse define a type of query
 * that is used internally by the datasource, to populate a progressive
 * control widget used to find GQL paths
 */
export interface FindChildrenRequest extends DataQuery {
  gql: string;
}

export interface FindChildrenResponse extends DataQueryResponse {
  nodes: GqlNode[];
}

/*
 * The JSON configuration object for this datasource.  This
 * along with the 'secret' data gets stored in Grafana.
 */
export interface AlcDataSourceOptionsJSON extends DataSourceJsonData {
  soapUrl: string;
  soapUser: string;
  allowPv: boolean;
  minPvCache: string;
}

export interface AlcDataSourceOptionsSECRET extends DataSourceJsonData {
  soapPassword?: string;
}

/*
 * This is the type for the list of query sub-editors for each individual
 * query type
 */
export type QueryViewGenerator = (props: AlcQueryProps, datasource: AlcDataSource, notify: QueryChangedHandler) => React.ReactNode;
export interface AlcQueryEditorDef {
  label: string;
  key: string;
  view: QueryViewGenerator;
}

export type QueryChangedHandler = (q: AlcQuerySpec) => void;
