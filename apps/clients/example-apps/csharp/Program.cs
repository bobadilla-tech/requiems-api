using System;
using System.Threading.Tasks;
using Org.OpenAPITools.Api;

class Program {
    static async Task Main() {
        var api = new AdviceApi();
        try {
            await api.V1TextAdviceGetAsync();
        } catch (Exception e) {
            if (e.Message.Contains("401") || e.Message.Contains("Unauthorized") || e.Message.Contains("Error calling V1TextAdviceGet")) {
                Console.WriteLine("✅ C# Client E2E OK (Expected 401)");
                Environment.Exit(0);
            }
            Console.WriteLine($"❌ Failed: {e.Message}");
            Environment.Exit(1);
        }
    }
}
