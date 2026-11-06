import { AdviceApi } from 'api-client';

async function run() {
  const api = new AdviceApi();
  try {
    await api.v1TextAdviceGet();
  } catch (e: any) {
    if (e.response && e.response.status === 401) {
      console.log('✅ TypeScript Client E2E OK (Expected 401)');
      process.exit(0);
    }
    console.error('❌ Failed:', e.message);
    process.exit(1);
  }
}
run();
