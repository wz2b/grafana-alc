import {
  CoreApp,
  DataSourceInstanceSettings,
  ScopedVars,
} from '@grafana/data';

import {
  DataSourceWithBackend,
  getTemplateSrv,
} from '@grafana/runtime';

import {
  AlcDataSourceOptionsJSON,
  AlcQueryProps,
} from './types';

const DEFAULT_QUERY: Partial<AlcQueryProps> = {
  alias: '',
  gql: '',
};

export class AlcDataSource extends DataSourceWithBackend<
    AlcQueryProps,
    AlcDataSourceOptionsJSON
> {
  constructor(
      instanceSettings: DataSourceInstanceSettings<AlcDataSourceOptionsJSON>
  ) {
    super(instanceSettings);
  }

  getDefaultQuery(_: CoreApp): Partial<AlcQueryProps> {
    return DEFAULT_QUERY;
  }

  applyTemplateVariables(
      query: AlcQueryProps,
      scopedVars: ScopedVars
  ): AlcQueryProps {
    return {
      ...query,
      gql: query.gql
          ? getTemplateSrv().replace(query.gql, scopedVars)
          : query.gql,
    };
  }

  filterQuery(query: AlcQueryProps): boolean {
    return !!query.gql;
  }
}

