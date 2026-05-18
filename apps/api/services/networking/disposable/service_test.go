package disposable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

func TestCheckEmail_Disposable(t *testing.T) {
	t.Parallel()
	svc := NewService()
	res := svc.CheckEmail("test@mailinator.com")
	assert.True(t, res.IsDisposable, "mailinator.com must be recognized as disposable")
	assert.Equal(t, "test@mailinator.com", res.Email)
	assert.Equal(t, "mailinator.com", res.Domain)
}

func TestCheckEmail_NotDisposable(t *testing.T) {
	t.Parallel()
	svc := NewService()
	res := svc.CheckEmail("user@gmail.com")
	assert.False(t, res.IsDisposable)
	assert.Equal(t, "gmail.com", res.Domain)
}

func TestCheckDomain_Disposable(t *testing.T) {
	t.Parallel()
	svc := NewService()
	res := svc.CheckDomain("mailinator.com")
	assert.True(t, res.IsDisposable)
	assert.Equal(t, "mailinator.com", res.Domain)
}

func TestCheckDomain_NotDisposable(t *testing.T) {
	t.Parallel()
	svc := NewService()
	res := svc.CheckDomain("gmail.com")
	assert.False(t, res.IsDisposable)
}

func TestCheckBatch_CountMatchesInput(t *testing.T) {
	t.Parallel()
	svc := NewService()
	emails := []string{"a@mailinator.com", "b@gmail.com", "c@mailinator.com"}
	res := svc.CheckBatch(emails)
	assert.Equal(t, 3, len(res))
	assert.Len(t, res, 3)
	assert.True(t, res[0].IsDisposable)
	assert.False(t, res[1].IsDisposable)
	assert.True(t, res[2].IsDisposable)
}

func TestGetStats_TotalNonZero(t *testing.T) {
	t.Parallel()
	svc := NewService()
	stats := svc.GetStats()
	assert.Greater(t, stats.TotalDomains, 0, "disposable domain list must be non-empty")
}

func TestGetDomains_PageLessThanOne(t *testing.T) {
	t.Parallel()
	svc := NewService()
	_, err := svc.GetDomains(0, 10)
	require.Error(t, err)
	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindInvalid, se.Kind)
}

func TestGetDomains_PerPageZero(t *testing.T) {
	t.Parallel()
	svc := NewService()
	_, err := svc.GetDomains(1, 0)
	require.Error(t, err)
	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindInvalid, se.Kind)
}

func TestGetDomains_PerPageTooLarge(t *testing.T) {
	t.Parallel()
	svc := NewService()
	_, err := svc.GetDomains(1, 1001)
	require.Error(t, err)
	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindInvalid, se.Kind)
}

func TestGetDomains_PageOutOfRange(t *testing.T) {
	t.Parallel()
	svc := NewService()
	_, err := svc.GetDomains(999999, 10)
	require.Error(t, err)
	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindNotFound, se.Kind)
	assert.Equal(t, "page_out_of_range", se.Code)
}

func TestGetDomains_HappyPath(t *testing.T) {
	t.Parallel()
	svc := NewService()
	res, err := svc.GetDomains(1, 5)
	require.NoError(t, err)
	assert.Equal(t, 1, res.Page)
	assert.Equal(t, 5, res.PerPage)
	assert.Len(t, res.Domains, 5)
	assert.Greater(t, res.Total, 0)
	assert.True(t, res.HasMore, "page 1 of 5 should have more pages")
}

func TestGetDomains_LastPage(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// Page large enough to be the last page.
	first, err := svc.GetDomains(1, 1000)
	require.NoError(t, err)
	totalPages := (first.Total + 999) / 1000

	last, err := svc.GetDomains(totalPages, 1000)
	require.NoError(t, err)
	assert.False(t, last.HasMore, "last page must not have more")
}

func TestExtractDomainFromEmail_Valid(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "example.com", ExtractDomainFromEmail("user@example.com"))
}

func TestExtractDomainFromEmail_NoAtSign(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", ExtractDomainFromEmail("notanemail"))
}

func TestExtractDomainFromEmail_MultipleAtSigns(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", ExtractDomainFromEmail("a@b@c"))
}
