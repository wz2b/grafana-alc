import React, { ChangeEvent } from 'react';
import {
  InlineField,
  Input,
  SecretInput,
  Switch,
} from '@grafana/ui';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';

import {
  AlcDataSourceOptionsJSON,
  AlcDataSourceOptionsSECRET,
} from '../types';

interface Props
    extends DataSourcePluginOptionsEditorProps<
        AlcDataSourceOptionsJSON,
        AlcDataSourceOptionsSECRET
    > {}

export function ConfigEditor(props: Props) {
  const { onOptionsChange, options } = props;

  const {
    jsonData,
    secureJsonFields,
    secureJsonData,
  } = options;

  const onSoapUrlChange = (event: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...jsonData,
        soapUrl: event.target.value,
      },
    });
  };

  const onSoapUserChange = (event: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...jsonData,
        soapUser: event.target.value,
      },
    });
  };

  const onMinPvCacheChange = (event: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...jsonData,
        minPvCache: event.target.value,
      },
    });
  };

  const onAllowPvChange = () => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...jsonData,
        allowPv: !jsonData.allowPv,
      },
    });
  };

  const onSoapPasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      secureJsonData: {
        ...secureJsonData,
        soapPassword: event.target.value,
      },
    });
  };

  const onResetSoapPassword = () => {
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...secureJsonFields,
        soapPassword: false,
      },
      secureJsonData: {
        ...secureJsonData,
        soapPassword: '',
      },
    });
  };

  return (
      <>
        <InlineField
            label="SOAP URL"
            labelWidth={18}
            tooltip="WebCTRL SOAP service URL"
        >
          <Input
              id="config-editor-soap-url"
              value={jsonData.soapUrl ?? ''}
              onChange={onSoapUrlChange}
              placeholder="https://webctrl.example.edu/ws"
              width={50}
          />
        </InlineField>

        <InlineField
            label="SOAP Username"
            labelWidth={18}
        >
          <Input
              id="config-editor-soap-user"
              value={jsonData.soapUser ?? ''}
              onChange={onSoapUserChange}
              width={40}
          />
        </InlineField>

        <InlineField
            label="SOAP Password"
            labelWidth={18}
        >
          <SecretInput
              id="config-editor-soap-password"
              isConfigured={Boolean(secureJsonFields?.soapPassword)}
              value={secureJsonData?.soapPassword ?? ''}
              onChange={onSoapPasswordChange}
              onReset={onResetSoapPassword}
              width={40}
          />
        </InlineField>

        <InlineField
            label="Allow Present Value"
            labelWidth={18}
        >
          <Switch
              value={jsonData.allowPv ?? false}
              onChange={onAllowPvChange}
          />
        </InlineField>

        <InlineField
            label="PV Cache Time"
            labelWidth={18}
            tooltip="Minimum cache duration for present-value requests, e.g. 30s or 1m"
        >
          <Input
              id="config-editor-min-pv-cache"
              value={jsonData.minPvCache ?? ''}
              onChange={onMinPvCacheChange}
              placeholder="1m"
              width={20}
          />
        </InlineField>
      </>
  );
}
