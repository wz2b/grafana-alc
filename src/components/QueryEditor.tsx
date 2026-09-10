import React, {ChangeEvent} from 'react';

import {QueryEditorProps, SelectableValue,} from '@grafana/data';

import {Combobox, ComboboxOption, InlineField, Input, Stack,} from '@grafana/ui';

import {AlcDataSource} from '../datasource';
import {AlcDataSourceOptionsJSON, AlcQueryProps,} from '../types';

type Props = QueryEditorProps<
    AlcDataSource,
    AlcQueryProps,
    AlcDataSourceOptionsJSON
>;

const queryTypeOptions: Array<ComboboxOption<string>> = [
    {
        label: 'Trend Data',
        value: 'Trend Data',
    },
    {
        label: 'Present Value',
        value: 'Present Value',
    },
    {
        label: 'Cache Stats',
        value: 'Cache Stats',
    },
];

export function QueryEditor({
                                query,
                                onChange,
                                onRunQuery,
                            }: Props) {
    const onQueryTypeChange = (option: SelectableValue<string>) => {
        onChange({
            ...query,
            queryType: option.value,
        });

        onRunQuery();
    };

    const onGqlChange = (event: ChangeEvent<HTMLInputElement>) => {
        onChange({
            ...query,
            gql: event.target.value,
        });
    };

    const onAliasChange = (event: ChangeEvent<HTMLInputElement>) => {
        onChange({
            ...query,
            alias: event.target.value,
        });
    };

    return (
        <Stack direction="column" gap={1}>
            <InlineField
                label="Query Type"
                labelWidth={16}
            >
                <Combobox
                    id="query-editor-query-type"
                    options={queryTypeOptions}
                    value={query.queryType}
                    onChange={onQueryTypeChange}
                    width={24}
                />
            </InlineField>

            {query.queryType !== 'Cache Stats' && (
                <InlineField
                    label="GQL"
                    labelWidth={16}
                    tooltip="WebCTRL GQL path"
                >
                    <Input
                        id="query-editor-gql"
                        value={query.gql ?? ''}
                        onChange={onGqlChange}
                        onBlur={onRunQuery}
                        placeholder="/Trees/geographic/..."
                        width={60}
                    />
                </InlineField>
            )}

            {query.queryType !== 'Cache Stats' && (
                <InlineField
                    label="Alias"
                    labelWidth={16}
                    tooltip="Optional display name for this query"
                >
                    <Input
                        id="query-editor-alias"
                        value={query.alias ?? ''}
                        onChange={onAliasChange}
                        onBlur={onRunQuery}
                        placeholder="Optional alias"
                        width={40}
                    />
                </InlineField>
            )}
        </Stack>
    );
}

