using System;
using Xunit;
using Moq;
using System.Text.Json;
using Microsoft.DevTunnels.Management;
using System.Text;

namespace Microsoft.DevTunnels.Test;

public class TunnelPlanTokenPropertiesTests
{
    private readonly string validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c";
    
    [Fact]
    public void TryParse_ReturnsNull_WhenTokenIsEmpty()
    {
        // Arrange
        string nullToken = null;
        string emptyToken = string.Empty;

        // Act
        Assert.Throws<ArgumentNullException>(() => TunnelPlanTokenProperties.TryParse(nullToken));
        Assert.Throws<ArgumentException>(() => TunnelPlanTokenProperties.TryParse(emptyToken));
    }

    [Fact]
    public void TryParse_ReturnsNull_WhenTokenHasInvalidFormat()
    {
        // Arrange
        string invalidToken1 = "invalidToken1";
        string invalidToken2 = "invalid.Token.2";
        string invalidToken3 = "invalid.Token.3";

        // Act
        var result1 = TunnelPlanTokenProperties.TryParse(invalidToken1);
        var result2 = TunnelPlanTokenProperties.TryParse(invalidToken2);
        var result3 = TunnelPlanTokenProperties.TryParse(invalidToken3);

        // Assert
        Assert.Null(result1);
        Assert.Null(result2);
        Assert.Null(result3);
    }

    [Fact]
    public void TryParse_ReturnsValue_WhenTokenIsValid()
    {
        // Act
        var result = TunnelPlanTokenProperties.TryParse(validToken);

        // Assert
        Assert.NotNull(result);
    }
}