import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
	schema: '../internal/transport/graphql/schema/**/*.graphqls',
	documents: ['src/lib/api/operations/**/*.graphql'],
	generates: {
		'./src/lib/api/generated/': {
			preset: 'client',
			config: {
				useTypeImports: true,
				enumsAsTypes: true,
				skipTypename: false,
				nonOptionalTypename: true,
				avoidOptionals: {
					field: false,
					inputValue: false,
					object: false
				}
			}
		}
	},
	ignoreNoDocuments: false
};

export default config;
