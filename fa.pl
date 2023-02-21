#!/usr/bin/env perl
use 5.020_000;
use warnings;
use Mojo::UserAgent qw<>;
use autodie;

my $page = '';
{
    my $res = Mojo::UserAgent->new->get('https://fontawesome.com/v4/cheatsheet/')->result;
    if ($res->is_success) {
        $page = $res->body;
    }
    elsif ($res->is_error) {
        die $res->message;
    }
    else {
        die "AIEE";
    }
}
# Example block:
#    <div class="col-md-4 col-sm-6 col-lg-3 col-print-4">
#      <small class="text-muted pull-right">4.7</small>
#      <i class="fa fa-fw" aria-hidden="true" title="Copy to use microchip">&#xf2db;</i>
#      fa-microchip
#      
#      <span class="text-muted">[&amp;#xf2db;]</span>
#    </div>
my $RX_BLOCK = qr{
    <div \s+ class="[^"]+">
        \s*
        (?:
            <small \s+ class="[^"]+">\s*(?<from_version>[^<]+)\s*</small>
        )?
        \s*
        <i[^>]+>[&][#]x(?<hex>[1-9a-fA-F][0-9a-fA-F]*);</i>
        \s*
        (?<name>\S+)
        \s*
        (?:
            <span[^>]+?>\[.+?\]</span>
        )?
        \s*
    </div>
}xms;
open my $fh, '>', 'fa.go';
printf {$fh} "package main\n\n";
printf {$fh} "type faRune struct {\n";
printf {$fh} "\tName        string\n";
printf {$fh} "\tFromVersion string\n";
printf {$fh} "}\n\n";
printf {$fh} "var faRunes = map[rune]faRune{\n";
while ($page =~ m!$RX_BLOCK!xmsg) {
    my ($from_version, $hex, $name) = ($+{from_version}, $+{hex}, $+{name});
    printf {$fh} sprintf qq!\t%-7s{"%s", "%s"},\n!,
        (sprintf '%d:', hex $hex),
        $name,
        $from_version // '';
}
printf {$fh} "}\n";
close $fh;
