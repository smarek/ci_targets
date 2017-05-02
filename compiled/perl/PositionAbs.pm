# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

use strict;
use warnings;
use IO::KaitaiStruct 0.007_000;
use Encode;

########################################################################
package PositionAbs;

our @ISA = 'IO::KaitaiStruct::Struct';

sub from_file {
    my ($class, $filename) = @_;
    my $fd;

    open($fd, '<', $filename) or return undef;
    binmode($fd);
    return new($class, IO::KaitaiStruct::Stream->new($fd));
}

sub new {
    my ($class, $_io, $_parent, $_root) = @_;
    my $self = IO::KaitaiStruct::Struct->new($_io);

    bless $self, $class;
    $self->{_parent} = $_parent;
    $self->{_root} = $_root || $self;

    $self->{index_offset} = $self->{_io}->read_u4le();

    return $self;
}

sub index {
    my ($self) = @_;
    return $self->{index} if ($self->{index});
    my $_pos = $self->{_io}->pos();
    $self->{_io}->seek($self->index_offset());
    $self->{index} = PositionAbs::IndexObj->new($self->{_io}, $self, $self->{_root});
    $self->{_io}->seek($_pos);
    return $self->{index};
}

sub index_offset {
    my ($self) = @_;
    return $self->{index_offset};
}

########################################################################
package PositionAbs::IndexObj;

our @ISA = 'IO::KaitaiStruct::Struct';

sub from_file {
    my ($class, $filename) = @_;
    my $fd;

    open($fd, '<', $filename) or return undef;
    binmode($fd);
    return new($class, IO::KaitaiStruct::Stream->new($fd));
}

sub new {
    my ($class, $_io, $_parent, $_root) = @_;
    my $self = IO::KaitaiStruct::Struct->new($_io);

    bless $self, $class;
    $self->{_parent} = $_parent;
    $self->{_root} = $_root || $self;

    $self->{entry} = Encode::decode("UTF-8", $self->{_io}->read_bytes_term(0, 0, 1, 1));

    return $self;
}

sub entry {
    my ($self) = @_;
    return $self->{entry};
}

1;
