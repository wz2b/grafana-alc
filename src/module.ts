import { DataSourcePlugin } from '@grafana/data';
import { AlcDataSource } from './datasource';
import { ConfigEditor } from './components/ConfigEditor';
import { QueryEditor } from './components/QueryEditor';
import { AlcQueryProps, AlcDataSourceOptionsJSON } from './types';

export const plugin = new DataSourcePlugin<AlcDataSource, AlcQueryProps, AlcDataSourceOptionsJSON>(AlcDataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
