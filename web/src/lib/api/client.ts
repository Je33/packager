import { env } from '$env/dynamic/public';
import type { TypedDocumentNode } from '@graphql-typed-document-node/core';
import { print } from 'graphql';

export class GraphQLClient {
	private url: string;

	constructor(url?: string) {
		this.url = url || env.PUBLIC_API_URL || 'http://localhost:8080/query';
	}

	/**
	 * Execute a type-safe GraphQL request
	 * @param document - Typed document node from generated code
	 * @param variables - Type-safe variables
	 * @returns Type-safe result
	 */
	async request<TResult, TVariables>(
		document: TypedDocumentNode<TResult, TVariables>,
		...[variables]: TVariables extends Record<string, never> ? [] : [TVariables]
	): Promise<TResult> {
		const query = print(document);

		const response = await fetch(this.url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				query,
				variables: variables ?? undefined
			})
		});

		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}

		const { data, errors } = await response.json();

		if (errors && errors.length > 0) {
			throw new GraphQLClientError(errors);
		}

		if (!data) {
			throw new Error('No data returned from GraphQL query');
		}

		return data;
	}
}

export class GraphQLClientError extends Error {
	errors: Array<{
		message: string;
		locations?: Array<{ line: number; column: number }>;
		path?: string[];
	}>;

	constructor(
		errors: Array<{
			message: string;
			locations?: Array<{ line: number; column: number }>;
			path?: string[];
		}>
	) {
		super(errors[0]?.message || 'GraphQL Error');
		this.name = 'GraphQLClientError';
		this.errors = errors;
	}
}

// Export singleton instance
export const graphqlClient = new GraphQLClient();
