package htmltest

import (
	"github.com/wjdp/htmltest/issues"
	"testing"
)

// Spec tests

func TestAnchorMissingHref(t *testing.T) {
	// fails for link with no href
	hT := t_testFile("fixtures/links/missingLinkHref.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "href blank", 1)
}

func TestAnchorIgnorable(t *testing.T) {
	// ignores links marked as ignore data-proofer-ignore
	hT := t_testFile("fixtures/links/ignorableLinks.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalLinkBroken(t *testing.T) {
	// fails for broken external links
	hT := t_testFile("fixtures/links/brokenLinkExternal.html")
	t_expectIssueCount(t, hT, 1)
}

func TestExternalLinkIgnore(t *testing.T) {
	// ignores external links when asked
	hT := t_testFileOpts("fixtures/links/brokenLinkExternal.html",
		map[string]interface{}{"CheckExternal": false})
	t_expectIssueCount(t, hT, 0)
}

func TestExternalHashBrokenDefault(t *testing.T) {
	// passes for broken external hashes by default
	hT := t_testFile("fixtures/links/brokenHashOnTheWeb.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalHashBrokenOption(t *testing.T) {
	// fails for broken external hashes when asked
	t.Skip("Not yet implemented")
	hT := t_testFile("fixtures/links/brokenHashOnTheWeb.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "no such hash", 1)
}

func TestExternalCache(t *testing.T) {
	// does not check links with parameters multiple times
	// TODO check cache is being checked
	t.Skip("Not yet implemented")
	hT := t_testFile("fixtures/links/check_just_once.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalHrefMalformed(t *testing.T) {
	// does not explode on bad external links in files
	hT := t_testFile("fixtures/links/bad_external_links.html")
	t_expectIssueCount(t, hT, 2)
}

func TestExternalInsecureDefault(t *testing.T) {
	// passes for non-HTTPS links when not asked
	hT := t_testFile("fixtures/links/non_https.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalInsecureOption(t *testing.T) {
	// fails for non-HTTPS links when asked
	hT := t_testFileOpts("fixtures/links/non_https.html",
		map[string]interface{}{"EnforceHTTPS": true})
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "is not an HTTPS target", 1)
}

func TestExternalHrefIP(t *testing.T) {
	// fails for broken IP address links
	hT := t_testFile("fixtures/links/ip_href.html")
	t_expectIssueCount(t, hT, 2)
}

func TestExternalHrefIPTimeout(t *testing.T) {
	// fails for broken IP address links
	hT := t_testFile("fixtures/links/ip_timeout.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "request exceeded our ExternalTimeout", 1)
}

func TestExternalFollowRedirects(t *testing.T) {
	// should follow redirects
	t.Skip("Need new link, times out")
	hT := t_testFile("fixtures/links/linkWithRedirect.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalFollowRedirectsDisabled(t *testing.T) {
	// fails on redirects if not following
	t.Skip("Not yet implemented, need new link, times out")
	hT := t_testFile("fixtures/links/linkWithRedirect.html")
	t_expectIssueCount(t, hT, 99)
	t_expectIssue(t, hT, "PLACEHOLDER", 99)
}

func TestExternalHTTPS(t *testing.T) {
	// should understand https
	hT := t_testFile("fixtures/links/linkWithHttps.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalMissingProtocolValid(t *testing.T) {
	// works for valid links missing the protocol
	hT := t_testFile("fixtures/links/link_missing_protocol_valid.html")
	t_expectIssueCount(t, hT, 0)
}

func TestExternalMissingProtocolInvalid(t *testing.T) {
	// fails for invalid links missing the protocol
	hT := t_testFile("fixtures/links/link_missing_protocol_invalid.html")
	t_expectIssueCount(t, hT, 1)
	// t_expectIssue(t, hT, "no such host", 1)
}

func TestExternalHrefPipes(t *testing.T) {
	// works for pipes in the URL
	hT := t_testFile("fixtures/links/escape_pipes.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalBroken(t *testing.T) {
	// fails for broken internal links
	hT := t_testFile("fixtures/links/brokenLinkInternal.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "target does not exist", 1)
}

func TestInternalRelativeLinksBase(t *testing.T) {
	// passes for relative links with a base
	t.Skip("Broken, ones does not exist, third back operation to base not supported")
	hT := t_testFile("fixtures/links/relativeLinksWithBase.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalHashBroken(t *testing.T) {
	// fails for broken internal hash
	t.Skip("Not yet implemented")
	hT := t_testFile("fixtures/links/brokenHashInternal.html")
	t_expectIssueCount(t, hT, 99)
	t_expectIssue(t, hT, "PLACEHOLDER", 99)
}

func TestDirectoryRootResolve(t *testing.T) {
	// properly resolves implicit /index.html in link paths
	hT := t_testFile("fixtures/links/linkToFolder.html")
	t_expectIssueCount(t, hT, 0)
}

func TestDirectoryCustomRoot(t *testing.T) {
	// works for custom directory index file
	t.Skip("Not yet implemented")
	hT := t_testFile("fixtures/links/link_pointing_to_directory.html")
	t_expectIssueCount(t, hT, 0)
}

func TestDirectoryCustomRootBroken(t *testing.T) {
	// fails if custom directory index file doesn't exist
	hT := t_testFile("fixtures/links/link_pointing_to_directory.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "target does not exist", 1)
}

func TestDirectoryNoTrailingSlash(t *testing.T) {
	// fails for internal linking to a directory without trailing slash
	hT := t_testFile("fixtures/links/link_directory_without_slash.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "target is a directory, href lacks trailing slash", 1)
}

func TestDirectoryHtmlExtension(t *testing.T) {
	// works for custom directory index file
	hT := t_testDirectory("fixtures/links/_site/")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalRootLink(t *testing.T) {
	// properly checks links to root
	hT := t_testFile("fixtures/links/rootLink/rootLink.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalRelativeLinks(t *testing.T) {
	// properly checks relative links
	hT := t_testFile("fixtures/links/relativeLinks.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalHrefNonstandardChars(t *testing.T) {
	// passes non-standard characters
	hT := t_testFile("fixtures/links/non_standard_characters.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalHrefUTF8(t *testing.T) {
	// passes for external UTF-8 links
	hT := t_testFile("fixtures/links/utf8Link.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalHrefUrlEncoded(t *testing.T) {
	// passes for urlencoded href
	hT := t_testFile("fixtures/links/urlencoded-href.html")
	t_expectIssueCount(t, hT, 0)
}

func TestErrorDuplication(t *testing.T) {
	// does not dupe errors
	hT := t_testFile("fixtures/links/nodupe.html")
	t_expectIssueCount(t, hT, 1)
}

func TestInternalDashedAttrs(t *testing.T) {
	// does not complain for files with attributes containing dashes
	hT := t_testFile("fixtures/links/attributeWithDash.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalCaseMismatch(t *testing.T) {
	// does not complain for internal links with mismatched cases
	hT := t_testFile("fixtures/links/ignores_cases.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalHashDefault(t *testing.T) {
	// fails for # href when not asked
	hT := t_testFile("fixtures/links/hash_href.html")
	t_expectIssue(t, hT, "empty hash", 1)
	t_expectIssueCount(t, hT, 1)
}

func TestInternalHashOption(t *testing.T) {
	// passes for # href when asked
	t.Skip("Not yet implemented")
	hT := t_testFile("fixtures/links/hash_href.html")
	t_expectIssueCount(t, hT, 0)
}

func TestInternalHashWeird(t *testing.T) {
	// works for internal links to weird encoding IDs
	hT := t_testFile("fixtures/links/encodingLink.html")
	t_expectIssueCount(t, hT, 0)
}

func TestMultipleProblems(t *testing.T) {
	// finds a mix of broken and unbroken links
	t.Skip("Only single problem, and an hash which is not yet supported.")
	// TODO make our own multiple problem file
	hT := t_testFile("fixtures/links/multipleProblems.html")
	t_expectIssueCount(t, hT, 99)
	t_expectIssue(t, hT, "PLACEHOLDER", 99)
}

func TestMailtoValid(t *testing.T) {
	// ignores valid mailto links
	hT := t_testFile("fixtures/links/mailto_link.html")
	t_expectIssueCount(t, hT, 0)
}

func TestMailtoBlank(t *testing.T) {
	// fails for blank mailto links
	hT := t_testFile("fixtures/links/blank_mailto_link.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "mailto is empty", 1)
}

func TestMailtoInvalid(t *testing.T) {
	// fails for invalid mailto links
	hT := t_testFile("fixtures/links/invalid_mailto_link.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "contains an invalid email address", 1)
}

func TestMailtoIgnore(t *testing.T) {
	// ignores mailto links when told to
	hT := t_testFileOpts("fixtures/links/blank_mailto_link.html",
		map[string]interface{}{"CheckMailto": false})
	t_expectIssueCount(t, hT, 0)
}

func TestTelValid(t *testing.T) {
	// ignores valid tel links
	hT := t_testFile("fixtures/links/tel_link.html")
	t_expectIssueCount(t, hT, 0)
}

func TestTelBlank(t *testing.T) {
	// fails for blank tel links
	hT := t_testFile("fixtures/links/blank_tel_link.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "tel is empty", 1)
}

func TestJavascriptLinkIgnore(t *testing.T) {
	// ignores javascript links
	hT := t_testFile("fixtures/links/javascript_link.html")
	t_expectIssueCount(t, hT, 0)
}

func TestLinkHrefValid(t *testing.T) {
	// works for valid href within link elements
	hT := t_testFile("fixtures/links/head_link_href.html")
	t_expectIssueCount(t, hT, 0)
}

func TestLinkHrefBlank(t *testing.T) {
	// fails for empty href within link elements
	hT := t_testFile("fixtures/links/head_link_href_empty.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "href blank", 1)
}

func TestLinkHrefAbsent(t *testing.T) {
	// fails for absent href within link elements
	hT := t_testFile("fixtures/links/head_link_href_absent.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "link tag missing href", 1)
}

// TODO invalid link href?

func TestAnchorPre(t *testing.T) {
	// works for broken anchors within pre & code
	hT := t_testFile("fixtures/links/anchors_in_pre.html")
	t_expectIssueCount(t, hT, 0)
}

func TestLinkPre(t *testing.T) {
	// works for broken link within pre & code
	hT := t_testFile("fixtures/links/links_in_pre.html")
	t_expectIssueCount(t, hT, 0)
}

func TestHashQueryBroken(t *testing.T) {
	// fails for broken hash with query
	t.Skip("Not yet dealt with")
	hT := t_testFile("fixtures/links/broken_hash_with_query.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "PLACEHOLDER", 99)
}

func TestHashSelf(t *testing.T) {
	// works for hash referring to itself
	hT := t_testFile("fixtures/links/hashReferringToSelf.html")
	t_expectIssueCount(t, hT, 0)
}

func TestAnchorNameIgnore(t *testing.T) {
	// ignores placeholder with name
	hT := t_testFile("fixtures/links/placeholder_with_name.html")
	t_expectIssueCount(t, hT, 0)
}

func TestAnchorIdIgnore(t *testing.T) {
	// ignores placeholder with id
	hT := t_testFile("fixtures/links/placeholder_with_id.html")
	t_expectIssueCount(t, hT, 0)
}

func TestAnchorIdEmpty(t *testing.T) {
	// fails for placeholder with empty id
	// TODO: Should we only fail here if missing href?
	t.Skip("Not yet implemented")
	hT := t_testFile("fixtures/links/placeholder_with_empty_id.html")
	t_expectIssueCount(t, hT, 1)
	t_expectIssue(t, hT, "anchor with empty id", 99)
}

func TestOtherProtocols(t *testing.T) {
	// ignores non-hypertext protocols
	hT := t_testFile("fixtures/links/other_protocols.html")
	t_expectIssueCount(t, hT, 0)
}

func TestAnchorBlankHTML5(t *testing.T) {
	// does not expect href for anchors in HTML5
	hT := t_testFile("fixtures/links/blank_href_html5.html")
	t_expectIssueCount(t, hT, 0)
}

func TestAnchorBlankHTML4(t *testing.T) {
	// does expect href for anchors in non-HTML5
	t.Skip("Not yet implemented")
	hT1 := t_testFile("fixtures/links/blank_href_html4.html")
	t_expectIssueCount(t, hT1, 1)
	hT2 := t_testFile("fixtures/links/blank_href_htmlunknown.html")
	t_expectIssueCount(t, hT2, 1)
}

func TestHTML5Page(t *testing.T) {
	// Page containing HTML5 tags
	hT := t_testFile("fixtures/html/html5_tags.html")
	t_expectIssueCount(t, hT, 0)
}

// TODO test canonical links
// TODO test "Unhandled client error"
// TODO test CheckInternal = false
// TODO test CheckExternal = false
// TODO test CheckMailto = false
// TODO test CheckTel = false

// Benchmarks

func BenchmarkManyExternalLinks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t_testFileOpts("fixtures/benchmarks/manyExternalLinks.html",
			map[string]interface{}{"LogLevel": issues.NONE})
	}
}

func BenchmarkManyExternalLinksDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t_testFileOpts("fixtures/benchmarks/manyExternalLinks.html",
			map[string]interface{}{"LogLevel": issues.NONE})
	}
}
