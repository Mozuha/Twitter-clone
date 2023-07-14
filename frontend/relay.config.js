module.exports = {
  src: './',
  language: 'typescript',
  schema: './schema/schema.graphql',
  artifactDirectory: './app/relay/__generated__',
  exclude: ['**/node_modules/**', '**/__mocks__/**', '**/__generated__/**'],
};
